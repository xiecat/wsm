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

func (r *ReqClient) DoRequest(url string, data string) (*Result, error) {
	resp, err := r.sendPayload(url, data)
	if err != nil {
		return nil, err
	}
	result, err := r.parseHttpResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ReqClient) sendPayload(u string, data string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, u, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	for k, v := range r.header {
		request.Header.Set(k, v)
	}
	resp, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReqClient) parseHttpResponse(resp *http.Response) (*Result, error) {
	defer resp.Body.Close()
	result := new(Result)
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	raw = bytes.TrimRight(raw, "\r\n\r\n")
	result.RawBody = raw
	result.Status = resp.StatusCode
	return result, nil
}
