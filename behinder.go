package wsm

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Go0p/wsm/lib/dynamic"
	"github.com/Go0p/wsm/lib/httpx"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/behinder"
	"github.com/Go0p/wsm/lib/utils"
)

type BehinderInfo struct {
	shell.BaseShell
	secretKey   []byte
	encryptMode int
	// response body 中的起始位
	beginIndex int
	// response body 中的结束位
	endIndex int
}

func NewBehinder(b BehinderInfo) *BehinderInfo {
	b.secretKey = utils.SecretKey(b.Password)
	return &b
}

// processParams 不同的方法需要传递不同的参数
func (b *BehinderInfo) processParams(p map[string]string) {
	if b.Script != shell.JavaScript {
		delete(p, "forcePrint")
		delete(p, "notEncrypt")
	}
}

func (b *BehinderInfo) Ping(p shell.IParams) bool {
	p.Check()
	params, err := utils.ToMapParams(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	b.processParams(params)
	data := behinder.GetData(b.secretKey, "EchoGo", params, b.Script, b.encryptMode)
	resultObj, ok := httpx.RequestAndParse(b.Url, b.Proxy, b.Headers, data, b.beginIndex, b.endIndex)
	if !ok {
		return false
	}
	localResultTxt := fmt.Sprintf(`{"msg":"%s","Status":"c3VjY2Vzcw=="}`, base64.StdEncoding.EncodeToString([]byte(params["content"])))
	localResultTxt2 := fmt.Sprintf(`{"Status":"c3VjY2Vzcw==","msg":"%s",}`, base64.StdEncoding.EncodeToString([]byte(params["content"])))
	var localResult, localResult2 []byte
	if params["notEncrypt"] == "true" {
		localResult = []byte(localResultTxt)
		localResult2 = []byte(localResultTxt2)
	} else {
		localResult = behinder.Encrypto([]byte(localResultTxt), b.secretKey, b.encryptMode, b.Script)
		localResult2 = behinder.Encrypto([]byte(localResultTxt2), b.secretKey, b.encryptMode, b.Script)
	}
	resData := resultObj.Data
	fmt.Println("resData", base64.StdEncoding.EncodeToString(resData))
	fmt.Println("localResult", base64.StdEncoding.EncodeToString(localResult))
	s1 := dynamic.MatchData(resData, localResult)
	s2 := dynamic.MatchData(resData, localResult2)
	if s1 < 0 && s2 < 0 {
		b.beginIndex = 0
		b.endIndex = 0
	} else if s1 >= 0 {
		b.endIndex = len(resData) - b.beginIndex - len(localResult)
	} else if s2 >= 0 {
		b.endIndex = len(resData) - b.beginIndex - len(localResult2)
	}
	resultTxt := behinder.Decrypto(resData[b.beginIndex:len(resData)-b.endIndex], b.secretKey, b.encryptMode, b.Script, params["notEncrypt"])
	fmt.Println("resultTxt", base64.StdEncoding.EncodeToString(resultTxt))
	result := make(map[string]string, 2)
	if err := json.Unmarshal(resultTxt, &result); err == nil {
		for k, v := range result {
			value, err := base64.StdEncoding.DecodeString(v)
			if err == nil {
				result[k] = string(value)
			}
		}
	}
	return false
}

func (b *BehinderInfo) BasicInfo() shell.Result {
	//b.Params.Check()
	//fmt.Printf("%#+v\n", b.Params)
	return nil
}
