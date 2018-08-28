package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
)

// 保存预约信息
type MicController struct {
	BaseController
}

//主页
func (this *MicController) Mic() {
	this.TplName = "mic.html"
}

func (this *MicController) SaveAppointment() {
	url := this.GetString("url")
	datas := this.GetString("datas")
	fmt.Println("接口地址=", url)
	fmt.Println("接口参数=", datas)

	response := HttpPost(url, datas)
	this.ajaxList("1", "成功", response)
}

func HttpPost(url string, datas string) string {

	//	resp, err := http.Post(url,
	//		"application/json",
	//		strings.NewReader("datas="+datas))

	//	reader := bytes.NewReader(datas)
	//	url := "http://localhost/echo.php"
	//	request, err := http.NewRequest("POST", url, reader)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	resp.Header.Set("Content-Type", "application/json;charset=UTF-8")

	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)

	//	fmt.Println(string(body))
	//	return string(body)

	var jsonStr = []byte(datas)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return string(body)
}
