package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dddengyunjie/heroes-service/common/pkgMarble"
	"github.com/dddengyunjie/heroes-service/web/common/app"
	"strconv"
)

type MarbleController struct {
	App *app.Application
	beego.Controller
}

func (this *MarbleController) InitMarble() {
	name := this.Ctx.Input.Param(":name")
	color := this.Ctx.Input.Param(":color")
	size, err := strconv.Atoi(this.Ctx.Input.Param(":size"))
	if err != nil {
		fmt.Println("Failed to initMarble:", err)
		return
	}
	owner := this.Ctx.Input.Param(":owner")
	var marble pkgMarble.Marble
	marble.Name = name
	marble.Color = color
	marble.Size = size
	marble.Owner = owner
	_, err = this.App.Fabric.InitMarble(marble)
	if err != nil {
		fmt.Println("Failed to initMarble:", err)
		this.Ctx.Output.Body([]byte("InitMarble failed"))
	} else {
		this.Ctx.Output.Body([]byte("InitMarble success"))
	}
}

func (this *MarbleController) DeleteMarble() {
	name := this.Ctx.Input.Param(":name")
	_, err := this.App.Fabric.DeleteMarble(name)
	if err != nil {
		fmt.Println("Failed to deleteMarble:", err)
		this.Ctx.Output.Body([]byte("DeleteMarble failed"))
	} else {
		this.Ctx.Output.Body([]byte("DeleteMarble success"))
	}
}

func (this *MarbleController) ReadMarble() {
	name := this.Ctx.Input.Param(":name")
	result, err := this.App.Fabric.ReadMarble(name)
	if err != nil {
		fmt.Println("Failed to ReadMarble:", err)
		this.Ctx.Output.Body([]byte("ReadMarble failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}

}

func (this *MarbleController) TransferMarble() {
	name := this.Ctx.Input.Param(":name")
	newOwner := this.Ctx.Input.Param(":newOwner")
	_, err := this.App.Fabric.TransferMarble(name, newOwner)
	if err != nil {
		fmt.Println("Failed to TransferMarble:", err)
		this.Ctx.Output.Body([]byte("TransferMarble failed"))
	} else {
		this.Ctx.Output.Body([]byte("TransferMarble success"))
	}
}

func (this *MarbleController) TransferMarbleBasedOnColor() {
	color := this.Ctx.Input.Param(":color")
	newOwner := this.Ctx.Input.Param(":newOwner")
	_, err := this.App.Fabric.TransferMarbleBasedOnColor(color, newOwner)
	if err != nil {
		fmt.Println("Failed to TransferMarbleBasedOnColor:", err)
		this.Ctx.Output.Body([]byte("TransferMarbleBasedOnColor failed"))
	} else {
		this.Ctx.Output.Body([]byte("TransferMarbleBasedOnColor success"))
	}
}

func (this *MarbleController) QueryMarblesByOwner() {
	owner := this.Ctx.Input.Param(":owner")
	result, err := this.App.Fabric.QueryMarblesByOwner(owner)
	if err != nil {
		fmt.Println("Failed to QueryMarblesByOwner:", err)
		this.Ctx.Output.Body([]byte("QueryMarblesByOwner failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}
}
