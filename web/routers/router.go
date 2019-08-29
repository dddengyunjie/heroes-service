package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//初始化 namespace
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/shop",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("shopinfo"))
				}),
			),
			beego.NSNamespace("/order",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("orderinfo"))
				}),
			),
			beego.NSNamespace("/crm",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("crminfo"))
				}),
			),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
