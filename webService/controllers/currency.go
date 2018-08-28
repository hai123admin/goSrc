package controllers

/**
*通用webService调用
*
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CurrencyController struct {
	BaseController
}

//主页
func (this *CurrencyController) IndexCurrency() {
	this.TplName = "currency.html"
}

//   webservice
func (this *CurrencyController) SendWebService() {
	var url string = this.GetString("url")

	var data = this.GetString("data")
	fmt.Println("")
	fmt.Println("webServie请求路径:", url)

	fmt.Println("webServie请求数据:", data)
	fmt.Println("")
	var result = PostWebServiceCurrency(url, data)
	this.ajaxList("1", "成功", result)
}

func PostWebServiceCurrency(url string, value string) string {

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

	return Byte2String(data)

}

func Byte2String(res []byte) string { return string(res) }

//   webservice
func (this *CurrencyController) PostQueueInfo() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("<<<<<<<<-----------推送排队信息------------>>>>>>>>>")

	var hospitalId string = this.GetString("hospitalId")

	fmt.Println("hospitalId=", hospitalId)

	var ob []QueueInfo //这是一个model，struct类型
	//var ob interface{}                 //这是一个model，struct类型
	body := this.Ctx.Input.RequestBody //这是获取到的json二进制数据
	json.Unmarshal(body, &ob)          //解析二进制json，把结果放进ob中

	//fmt.Println(string(body)) //这是获取到的json

	var i int
	for i = 0; i < len(ob); i++ {
		fmt.Printf("%v\n", ob[i])
		//fmt.Printf("%+v\n", ob[i]) // {1 2}
	}

	objmsg := make(map[string]interface{})
	objmsg["IsSuccess"] = true
	objmsg["ResultCode"] = 100
	objmsg["ResultMsg"] = "测试成功"
	this.Data["json"] = objmsg
	this.ServeJSON()
	this.StopRun()

}

type QueueInfo struct {
	PatientName     string `json:"patientName"`
	ApplyDepartment string `json:"applyDepartment"`
	CheckItemName   string `json:"checkItemName"`
	VideoRoom       string `json:"videoRoom"`
	QueueNo         string `json:"queueNo"`
	Birthday        string `json:"birthday"`
	HisId           string `json:"hisId"`
	CheckNo         string `json:"checkNo"`
	AdmId           string `json:"admId"`
	AccessionNo     string `json:"accessionNo"`
	IsExigence      string `json:"isExigence"`
	CallFlag        string `json:"callFlag"`
	Sex             string `json:"sex"`
	ApplyParam      string `json:"applyParam"`
	ModalityName    string `json:"modalityName"`
	EnrolTime       string `json:"enrolTime"`
	Memo            string `json:"memo"`
}
