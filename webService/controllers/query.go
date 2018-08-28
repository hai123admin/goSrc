package controllers

/**
*无锡webService调用
*
 */
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	//"webService/models"

	"github.com/astaxie/beego/orm"
)

type QueryController struct {
	BaseController
}

//主页
func (this *QueryController) Get() {
	this.TplName = "index.html"
}

//主页
func (this *QueryController) Index() {
	this.TplName = "index.html"
}

func (this *QueryController) Main() {
	this.TplName = "main.html"
}

func (this *QueryController) AllDengji() {
	this.TplName = "alldengji.html"
}

func (this *QueryController) WuXiDengji() {
	this.TplName = "wuxidengji.html"
}

//主页
func (this *QueryController) Help() {
	this.TplName = "help.html"
}

//login
func (this *QueryController) Login() {
	this.TplName = "login.html"
}

//获取机构
func (this *QueryController) GetHospitalInfoList() {
	id := this.GetString("applyHospital")
	var sql string
	if id != "" {
		sql = "SELECT  HOSPITAL_PK as hospitalPk,HOSPITAL_NAME  as hospitalName  FROM hospital_info where  STATUS='A' AND HOSPITAL_PK = " + id + "  order by HOSPITAL_PK asc"
	} else {
		sql = "SELECT  HOSPITAL_PK as hospitalPk,HOSPITAL_NAME  as hospitalName  FROM hospital_info where  STATUS='A' order by HOSPITAL_PK asc"
	}
	//sql := "SELECT  HOSPITAL_PK as hospitalPk,HOSPITAL_NAME  as hospitalName  FROM hospital_info where  STATUS='A' AND HOSPITAL_PK = " + id + "  order by HOSPITAL_PK asc"
	var maps []orm.Params
	orm.NewOrm().Raw(sql).Values(&maps)
	this.ajaxList("1", "成功", maps)
}

//获取根据机构ID和应用类型查询科室
func (this *QueryController) GetDeptByHosAndTypeList() {
	id := this.GetString("applyHospital")
	ra := this.GetString("applicationType")

	sqlFlag := "SELECT  PARAM_1 FROM HOSPITAL_CONFIG WHERE HOSPITAL_INFO_FK =  " + id + " LIMIT 1 "
	var data []orm.Params
	orm.NewOrm().Raw(sqlFlag).Values(&data)
	var param1 string = "N"
	for _, tmp := range data {
		if tmp["PARAM_1"] != 'Y' {
			param1 = "Y"
		} else {
			param1 = "N"
		}

	}
	sql := "SELECT DEPARTMENT_PK AS departmentPk,	DEPARTMENT_CODE AS departmentCode,	DEPARTMENT_NAME AS departmentName,	DEPARTMENT_NOTE_NAME AS departmentNoteName FROM HOSPITAL_DEPARTMENT WHERE  IS_ENABLE = 'A' AND HOSPITAL_FK =" + id + " and APPLICATION_TYPE ='" + ra + "'"
	var maps []orm.Params
	orm.NewOrm().Raw(sql).Values(&maps)
	this.ajaxList("1", param1, maps)
}

//获取科室
func (this *QueryController) GetDeptentMentList() {
	id := this.GetString("applyHospital")
	applicationType := this.GetString("applicationType")
	sql := "SELECT DEPARTMENT_PK AS departmentPk,	DEPARTMENT_CODE AS departmentCode,	DEPARTMENT_NAME AS departmentName,	DEPARTMENT_NOTE_NAME AS departmentNoteName FROM hospital_department WHERE  IS_ENABLE = 'A' AND HOSPITAL_FK =" + id + "  AND APPLICATION_TYPE = '" + applicationType + "'"
	var maps []orm.Params
	orm.NewOrm().Raw(sql).Values(&maps)
	this.ajaxList("1", "成功", maps)
}

//获取检查类型
func (this *QueryController) GetModailtyList() {
	id := this.GetString("applyHospital")
	departmentName := this.GetString("departmentName")

	sql := "SELECT MODALITY_PK AS modalityPk,	MODALITY_CODE AS modalityCode,	MODALITY_NOTE_NAME AS modalityName  FROM hospital_modality WHERE HOSPITAL_FK =" + id + " AND IS_ENABLE = 'A' AND DEPARTMENT_FK =" + departmentName
	var maps []orm.Params
	orm.NewOrm().Raw(sql).Values(&maps)
	this.ajaxList("1", "成功", maps)
}

//获取检查项目
func (this *QueryController) GetCheckItemList() {
	id := this.GetString("applyHospital")
	fk := this.GetString("modalityName")

	sql := "SELECT 	ITEM_PK AS ItemPk,ITEM_CODE as ItemCode ,ITEM_NAME AS ItemName,ITEM_NOTE_NAME AS ItemNoteName,	CHECK_ITEM_FEE AS checkItemFee FROM hospital_check_item WHERE  HOSPITAL_FK=" + id + " AND MODALITY_FK =" + fk
	var maps []orm.Params
	orm.NewOrm().Raw(sql).Values(&maps)
	this.ajaxList("1", "成功", maps)
}

//获取影像室
func (this *QueryController) GetVideoRoomList() {
	id := this.GetString("applyHospital")
	fk := this.GetString("modalityName")

	sql := "SELECT VIDEO_ROOM_PK AS videoRoomPk,	VIDEO_ROOM_CODE AS videoRoomCode,	VIDEO_ROOM_NAME AS videoRoomName FROM hospital_video_room WHERE HOSPITAL_FK =" + id + " AND  MODALITY_FK = " + fk + " AND IS_ENABLE = 'A'"

	var maps []orm.Params
	orm.NewOrm().Raw(sql).Values(&maps)
	this.ajaxList("1", "成功", maps)
}

//  无锡 登记webservice
func (this *QueryController) InsertRegisterInfo() {
	var url string = this.GetString("url")

	var patientInfo = this.GetString("patientInfo")
	var checkApplyInfo = this.GetString("checkApplyInfo")
	var itemInfo = this.GetString("itemInfo")

	fmt.Println("")
	fmt.Println(patientInfo)
	fmt.Println(checkApplyInfo)
	fmt.Println(itemInfo)
	fmt.Println("")
	postStr := CreateSOAPXml(patientInfo, checkApplyInfo, itemInfo) // verify 为参数  http://itrus.com/itrusUtil为参数，参数内容根据实际情况获取
	var result = PostWebService(url, postStr)
	this.ajaxList("1", "成功", result)
}

func PostWebService(url string, value string) string {

	res, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(value)))
	//这里随便传递了点东西
	if err != nil {
		fmt.Println("post error", err)
	}
	data, err := ioutil.ReadAll(res.Body)

	//取出主体的内容
	if err != nil {
		fmt.Println("read error", err)
	}
	res.Body.Close()

	return ByteToString(data)

}

func ByteToString(res []byte) string { return string(res) }

//拼接参数
func CreateSOAPXml(patientInfo string, checkApplyInfo string, itemInfo string) string {

	soapBody := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.webservice.amol.accurad.com/">`
	soapBody += `<soapenv:Header/>`
	soapBody += `<soapenv:Body>`
	soapBody += `<ser:insertRegisterInfo>`

	// soapBody += `<patientInfo>{"patientName":"MR测试10","sex":"M","birthday":"1985-09-04","telephone":"18303517744","address":"陕西西安科技二路","idcardNo":"620621198011148999"}</patientInfo>`
	// soapBody += `<checkApplyInfo>{"hisId":"11225","admIdIss":"1","admId":"1","accessionNo":"11225","applyDepartmentStr":"急诊科","applyHospital":"无锡市人民医院","applyHospitalFk":"100651","modalityName":"MR","departmentName":"050105","applyParam":"0","videoRoomName":"G区磁共振室","applyCheckTimeStr":"2018-07-27 10:02:01","checkNo":"17000005","checkPurpose":"目的","applyDoctor":"张三","barcode":"15","isExigence":"Y","regionNo":"1","bedNo":"1-16","chiefComplaint":"诊断四大四大","diagnosis":"诊断","hisType":"11"}</checkApplyInfo>`
	// soapBody += `<itemInfo>[{"checkItemName":"MRI-心脏功能检查","extNo":"100","checkItemFee":"100"}]</itemInfo>`

	soapBody += "<patientInfo>" + patientInfo + "</patientInfo>"
	soapBody += "<checkApplyInfo>" + checkApplyInfo + "</checkApplyInfo>"
	soapBody += "<itemInfo>" + itemInfo + "</itemInfo>"

	soapBody += `</ser:insertRegisterInfo>`
	soapBody += `</soapenv:Body>`
	soapBody += `</soapenv:Envelope>`

	return soapBody
}
