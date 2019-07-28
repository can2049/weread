package controllers

import (
	"weread/models"
)

type GroupController struct {
	CommonController
}

// func (this *GroupController) Index() {
// 	if this.IsAjax() {
// 		page, _ := this.GetInt64("page")
// 		page_size, _ := this.GetInt64("rows")
// 		sort := this.GetString("sort")
// 		order := this.GetString("order")
// 		if len(order) > 0 {
// 			if order == "desc" {
// 				sort = "-" + sort
// 			}
// 		} else {
// 			sort = "Id"
// 		}
// 		nodes, count := models.GetGrouplist(page, page_size, sort)
// 		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
// 		this.ServeJSON()
// 		return
// 	} else {
// 		this.TplName = this.GetTemplatetype() + "/rbac/group.tpl"
// 	}
// }

func (this *GroupController) AddGroup() {
	g := models.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := models.AddGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) UpdateGroup() {
	g := models.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := models.UpdateGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) DelGroup() {
	Id, _ := this.GetInt64("Id")
	status, err := models.DelGroupById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
