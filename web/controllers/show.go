package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dddengyunjie/heroes-service/web/common/app"
)

type ShowController struct {
	App *app.Application
	beego.Controller
}

func (this *ShowController) Show() {
	helloValue, err := this.App.Fabric.QueryHello()
	if err != nil {
		this.Ctx.Output.Body([]byte("QueryHello error"))
	}
	this.Ctx.Output.Body([]byte("show helloValue:" + helloValue))
}
