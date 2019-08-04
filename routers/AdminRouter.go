package routers

import (
	"github.com/astaxie/beego"
	"weread/controllers"
)

func init() {
	beego.Router("/admin/login", &controllers.AdminController{}, "get:AdminLogin")
	beego.Info("Routers init finished")
}
