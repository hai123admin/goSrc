package routers

import (
	"webService/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// default
	beego.Router("/", &controllers.QueryController{})
	// index
	beego.Router("/index", &controllers.QueryController{}, "*:Index")
	// main
	beego.Router("/main", &controllers.QueryController{}, "*:Main")
	// other
	beego.Router("/help", &controllers.QueryController{}, "*:Help")

	beego.Router("/alldengji", &controllers.QueryController{}, "*:AllDengji")

	//无锡
	beego.Router("/wuxidengji", &controllers.QueryController{}, "*:WuXiDengji")

	beego.Router("/login", &controllers.QueryController{}, "*:Login")

	beego.Router("/getHospitalInfoList", &controllers.QueryController{}, "*:GetHospitalInfoList")

	beego.Router("/getDeptByHosAndTypeList", &controllers.QueryController{}, "*:GetDeptByHosAndTypeList")

	beego.Router("/getDeptentMentList", &controllers.QueryController{}, "*:GetDeptentMentList")

	beego.Router("/getModailtyList", &controllers.QueryController{}, "*:GetModailtyList")

	beego.Router("/getCheckItemList", &controllers.QueryController{}, "*:GetCheckItemList")

	beego.Router("/getVideoRoomList", &controllers.QueryController{}, "*:GetVideoRoomList")

	beego.Router("/insertRegisterInfo", &controllers.QueryController{}, "*:InsertRegisterInfo")

	beego.Router("/insertRegisterInfo", &controllers.QueryController{}, "*:InsertRegisterInfo")

	beego.Router("/indexCurrency", &controllers.CurrencyController{}, "*:IndexCurrency")

	beego.Router("/sendWebService", &controllers.CurrencyController{}, "*:SendWebService")

	// queue info
	beego.Router("/postQueueInfo", &controllers.CurrencyController{}, "*:PostQueueInfo")

	beego.Router("/mic", &controllers.MicController{}, "*:Mic")
	beego.Router("/saveAppointment", &controllers.MicController{}, "*:SaveAppointment")

	// error 400 500
	beego.ErrorController(&controllers.ErrorController{})

}
