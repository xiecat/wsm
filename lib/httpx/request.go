package httpx

import (
	"bytes"
	"crypto/tls"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/godzilla"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

const (
	DialTimout = 5 * time.Second
	KeepAlive  = 15 * time.Second
)

type Result struct {
	Data   []byte
	Status int
	Header http.Header
}

type ReqClient struct {
	client *http.Client
	header map[string]string
	script shell.ScriptType
	crypto godzilla.CrypticType
}

func NewClient(proxyUrl string, header map[string]string, script shell.ScriptType, crypt godzilla.CrypticType) *ReqClient {
	if len(header["User-Agent"]) == 0 {
		header = randUa(header)
	}
	return &ReqClient{
		client: newClient(proxyUrl),
		header: header,
		script: script,
		crypto: crypt,
	}
}

func randUa(header map[string]string) map[string]string {
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36"
	return header
}

func newClient(proxyUrl string) *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		log.Fatal(err)
	}
	dialer := &net.Dialer{
		Timeout:   DialTimout,
		KeepAlive: KeepAlive,
	}
	tr := &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConns:        1000,
		IdleConnTimeout:     KeepAlive,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 5 * time.Second,
	}
	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err == nil {
			// 使用传入代理
			tr.Proxy = http.ProxyURL(proxy)
		}
	}
	return &http.Client{
		Jar:       jar,
		Timeout:   time.Second * 25,
		Transport: tr,
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func (r *ReqClient) DoRequest(url string, data string, beginIndex, endIndex int) (*Result, bool) {
	resultObj := r.sendPayload(url, data)
	resData := bytes.TrimRight(resultObj.Data, "\r\n\r\n")
	if (beginIndex != 0 || endIndex != 0) && len(resData)-endIndex >= beginIndex {
		resData = resData[beginIndex : len(resData)-endIndex]
	}
	resultObj.Data = resData
	if len(resData) == 0 && resultObj.Status == 404 {
		return nil, false
	} else if len(resData) == 0 && resultObj.Status == 200 {
		return nil, true
	}
	return resultObj, true
}

func (r *ReqClient) setHeader() http.Header {
	h := make(http.Header, 2)
	if r.script == shell.JavaScript {
		if r.crypto == godzilla.JAVA_AES_BASE64 {
			h.Set("Content-type", "application/x-www-form-urlencoded")
		} else if r.crypto == godzilla.JAVA_AES_RAW {
		} else {
			log.Println("encryption mode err")
		}
	} else if r.script == shell.CsharpScript {
		if r.crypto == godzilla.CSHARP_AES_BASE64 {
		} else if r.crypto == godzilla.CSHARP_AES_RAW {
		} else {
			log.Println("encryption mode err")
		}
	} else if r.script == shell.PhpScript {
		if r.crypto == godzilla.PHP_XOR_BASE64 {
			h.Set("Content-type", "application/x-www-form-urlencoded")
		} else if r.crypto == godzilla.PHP_XOR_RAW {
		} else {
			log.Println("encryption mode err")
		}
	} else if r.script == shell.AspScript {
		if r.crypto == godzilla.ASP_XOR_BASE64 {
			h.Set("Content-type", "application/x-www-form-urlencoded")
		} else if r.crypto == godzilla.ASP_XOR_RAW {
		} else {
			log.Println("encryption mode err")
		}
	} else {
	}
	for k, v := range r.header {
		h.Set(k, v)
	}
	return h
}

func (r *ReqClient) sendPayload(u string, data string) *Result {
	result := new(Result)
	request, err := http.NewRequest(http.MethodPost, u, strings.NewReader(data))
	if err != nil {
		return nil
	}
	request.Header = r.setHeader()
	resp, err := r.client.Do(request)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result.Data = body
	result.Header = resp.Header
	result.Status = resp.StatusCode
	return result
}
