package controllers

import (
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
	Templatetype string //ui template type
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func (this *CommonController) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "easyui"
	}
	return templatetype
}

func init() {

	//验证权限
	AccessRegister()
}
