package httpx

import (
	"bytes"
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Result struct {
	Data   []byte
	Status int
	Header http.Header
}

var globalClient *http.Client

func init() {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		log.Fatal(err)
	}
	globalClient = &http.Client{
		Jar:     jar,
		Timeout: time.Second * 25,
	}
}

func RequestAndParse(urlPath, proxy string, header map[string]string, data string, beginIndex, endIndex int) (*Result, bool) {
	resultObj := sendPostRequestBinary(urlPath, proxy, header, data)
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

func sendPostRequestBinary(urlPath, proxyUrl string, headerMap map[string]string, data string) *Result {
	result := new(Result)
	request, err := http.NewRequest(http.MethodPost, urlPath, strings.NewReader(data))
	//request.Header.Set("Cookie", "JSESSIONID=C9E1460ACF024D2240E2C748DB54FA18;")
	for k, v := range headerMap {
		request.Header.Set(k, v)
	}
	if err != nil {
		return nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err == nil { // 使用传入代理
			tr.Proxy = http.ProxyURL(proxy)
		}
	}
	globalClient.Transport = tr

	resp, err := globalClient.Do(request)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result.Data = body
	result.Header = resp.Header
	result.Status = resp.StatusCode
	//fmt.Println("rs.Header:", resp.Header)
	return result
}
