package main

import (
	"github.com/astaxie/beego"
	"weread/models"
	_ "weread/routers"
	. "weread/utils"
	// "fmt"
	"os"
	// "mime"
)

func main() {
	initialize()
	beego.Run()
}

func initialize() {
	//mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()

	models.Connect()

	beego.AddFuncMap("stringsToJson", StringsToJson)
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}
