package main

import (
	"io/ioutil"
	"os"
	"webService/models"
	_ "webService/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//解析text文件内容
func LogoOut() {
	//打开文件的路径
	fi, err := os.Open("logo.txt")
	if err != nil {
		logs.Debug("打开文件失败")
	}
	defer fi.Close()
	//读取文件的内容
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		logs.Debug("读取文件失败")

	}
	var str = string(fd)
	logs.Debug(str)

}
func main() {
	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	//为了让日志输出不影响性能，开启异步日志
	logs.Async()
	//开启seesion支持，默认使用的存储引擎为：memory
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	models.Init()
	//beego.Register(&models.Init())
	LogoOut()
	logs.Debug("★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★")
	logs.Debug("                                                         ")
	logs.Debug("                                                         ")
	logs.Debug("              Yi Zhen Yun Web Service Test 启动           ")
	logs.Debug("                                                         ")
	logs.Debug("                                                         ")
	logs.Debug("★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★")

	//默认static目录是可以直接访问的，其它目录需要单独指定
	beego.SetStaticPath("/theme", "theme")

	//启动应用
	beego.Run()
}
