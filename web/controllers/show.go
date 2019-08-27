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
	helloValue, err := this.App.Fabric.QueryHello(false)
	if err != nil {
		this.Ctx.Output.Body([]byte("QueryHello error"))
	}
	this.Ctx.Output.Body([]byte("show helloValue:" + helloValue))
}

func (this *ShowController) ShowHistory() {
	helloValue, err := this.App.Fabric.QueryHello(true)
	if err != nil {
		this.Ctx.Output.Body([]byte("QueryHello error"))
	}
	this.Ctx.Output.Body([]byte("show helloValue:" + helloValue))
}
