package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dddengyunjie/heroes-service/web/common/app"
)

type UploadController struct {
	App *app.Application
	beego.Controller
}

func (this *UploadController) Upload() {
	value := this.Ctx.Input.Param(":value")
	_, err := this.App.Fabric.InvokeHello(value)
	if err != nil {
		this.Ctx.Output.Body([]byte("upload error"))
	}

	this.Ctx.Output.Body([]byte("upload something"))
}
