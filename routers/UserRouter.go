package routers

import (
	"github.com/astaxie/beego"
	"weread/controllers"
)

func init() {
	beego.Router("/", &controllers.UserController{}, "*:Index")
}
