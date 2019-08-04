package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
	"weread/models"
	_ "weread/routers"
	. "weread/utils"
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

	initSession()
	initTemplate()
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

func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.AppConfig.Session.SessionName = "liteblog-key"
	// beego.BConfig.WebConfig.Session.SessionProvider = "file"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
	beego.AddFuncMap("eq2", func(x, y interface{}) bool {
		s1 := fmt.Sprintf("%v", x)
		s2 := fmt.Sprintf("%v", y)
		return strings.Compare(s1, s2) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
	beego.AddFuncMap("json", func(obj interface{}) string {
		bs, err := json.Marshal(obj)
		if err != nil {
			return "{id:0}"
		}
		return string(bs)
	})
}
