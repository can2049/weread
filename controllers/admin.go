package controllers

import (
	// "errors"
	"github.com/astaxie/beego"
	// "github.com/satori/go.uuid"
	// "weread/models"
	// "liteblog/syserrors"
)

type AdminController struct {
	CommonController
	// beego.Controller
}

/*
func (ctx *CommonController) Prepare() {
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	// ctx.Dao = models.NewDB()
	// 验证用户是否登陆
	ctx.IsLogin = false
	if u, ok := ctx.GetSession("userinfo").(models.User); ok {
		// ctx.User = u
		ctx.Data["User"] = u
		ctx.IsLogin = true
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
	//判断子controller是否实现接口 NestPreparer
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
*/

func (this *AdminController) AdminLogin() {
	beego.Info("begin Admin login...")
	isajax := this.GetString("isajax")
	username := this.GetString("username")
	if username != "amdin" {
		this.Rsp(true, "not user: admin")
	}
		if isajax == "1" {
		password := this.GetString("password")
		user, err := CheckLogin(username, password)
		if err == nil {
			this.SetSession("userinfo", user)
			// accesslist, _ := GetAccessList(user.Id)
			// this.SetSession("accesslist", accesslist)
			this.Rsp(true, "登录成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}

	}
	userinfo := this.GetSession("userinfo")
	if userinfo != nil {
		this.Ctx.Redirect(302, "/admin/login")
	}
	this.TplName = "views/AdminIndex.html"
}
