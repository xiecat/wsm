package httpx

import (
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	DialTimout = 5 * time.Second
	KeepAlive  = 15 * time.Second
)

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
