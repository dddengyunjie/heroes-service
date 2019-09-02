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

func (this *MarbleController) TransferMarblesBasedOnColor() {
	color := this.Ctx.Input.Param(":color")
	newOwner := this.Ctx.Input.Param(":newOwner")
	_, err := this.App.Fabric.TransferMarblesBasedOnColor(color, newOwner)
	if err != nil {
		fmt.Println("Failed to TransferMarblesBasedOnColor:", err)
		this.Ctx.Output.Body([]byte("TransferMarblesBasedOnColor failed"))
	} else {
		this.Ctx.Output.Body([]byte("TransferMarblesBasedOnColor success"))
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

func (this *MarbleController) GetMarblesByRange() {
	startKey := this.Ctx.Input.Param(":startKey")
	endKey := this.Ctx.Input.Param(":endKey")
	result, err := this.App.Fabric.GetMarblesByRange(startKey, endKey)
	if err != nil {
		fmt.Println("Failed to GetMarblesByRange:", err)
		this.Ctx.Output.Body([]byte("GetMarblesByRange failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}
}

func (this *MarbleController) GetMarblesByRangeWithPagination() {
	startKey := this.Ctx.Input.Param(":startKey")
	endKey := this.Ctx.Input.Param(":endKey")
	pageSize, err := strconv.Atoi(this.Ctx.Input.Param(":pageSize"))
	bookMark := this.Ctx.Input.Param(":bookMark")
	if bookMark == "nil" { //FIXME 暂时用这个方法，需要完善
		bookMark = ""
	}
	if err != nil {
		fmt.Println("Failed to initMarble:", err)
		return
	}
	result, err := this.App.Fabric.GetMarblesByRangeWithPagination(startKey, endKey, pageSize, bookMark)
	if err != nil {
		fmt.Println("Failed to getMarblesByRangeWithPagination:", err)
		this.Ctx.Output.Body([]byte("getMarblesByRangeWithPagination failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}
}

func (this *MarbleController) QueryMarbles() {
	queryString := this.Ctx.Input.Param(":splat")
	result, err := this.App.Fabric.QueryMarbles(queryString)
	if err != nil {
		fmt.Println("Failed to QueryMarbles:", err)
		this.Ctx.Output.Body([]byte("QueryMarbles failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}
}

func (this *MarbleController) GetHistoryForMarble() {
	name := this.Ctx.Input.Param(":name")
	result, err := this.App.Fabric.GetHistoryForMarble(name)
	if err != nil {
		fmt.Println("Failed to GetHistoryForMarble:", err)
		this.Ctx.Output.Body([]byte("GetHistoryForMarble failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}
}

func (this *MarbleController) QueryMarblesWithPagination() {
	//queryString := this.Ctx.Input.Param(":queryString")
	pageSize, err := strconv.Atoi(this.Ctx.Input.Param(":pageSize"))
	bookMark := this.Ctx.Input.Param(":bookMark")
	if bookMark == "nil" { //FIXME 暂时用这个方法，需要完善
		bookMark = ""
	}
	queryString := this.Ctx.Input.Param(":splat")
	if err != nil {
		fmt.Println("Failed to initMarble:", err)
		return
	}
	result, err := this.App.Fabric.QueryMarblesWithPagination(queryString, pageSize, bookMark)
	if err != nil {
		fmt.Println("Failed to QueryMarblesWithPagination:", err)
		this.Ctx.Output.Body([]byte("QueryMarblesWithPagination failed"))
	} else {
		this.Ctx.Output.Body([]byte(result))
	}
}

//TestInitMarble
//测试需要，批量运行InitMarble函数
func (this *MarbleController) TestInitMarble() {
	name := this.Ctx.Input.Param(":name")
	color := this.Ctx.Input.Param(":color")
	size, err := strconv.Atoi(this.Ctx.Input.Param(":size"))
	if err != nil {
		fmt.Println("Failed to initMarble:", err)
		return
	}
	owner := this.Ctx.Input.Param(":owner")
	number, err := strconv.Atoi(this.Ctx.Input.Param(":number"))
	if err != nil {
		fmt.Println("Failed to initMarble:", err)
		return
	}

	var marble pkgMarble.Marble
	marble.Color = color
	marble.Size = size
	marble.Owner = owner
	for i := 0; i < number; i++ {
		marble.Name = name + strconv.Itoa(i)
		_, err := this.App.Fabric.InitMarble(marble)
		if err != nil {
			fmt.Println("Failed to initMarble:", err)
			this.Ctx.Output.Body([]byte("InitMarble failed"))
		} else {
			this.Ctx.Output.Body([]byte("InitMarble success"))
		}
	}
}
