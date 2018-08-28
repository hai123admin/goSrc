package controllers

/**
*基本BaseController
*
 */
import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

//前期准备
func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	fmt.Println(controllerName)
	fmt.Println(actionName)

	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	fmt.Println(this.controllerName)
	fmt.Println(this.actionName)
}

//ajax返回 列表
func (this *BaseController) ajaxList(msg string, code string, data interface{}) {

	objmsg := make(map[string]interface{})
	objmsg["code"] = code
	objmsg["msg"] = msg
	objmsg["data"] = data
	this.Data["json"] = objmsg

	this.ServeJSON()

	this.StopRun()
}
