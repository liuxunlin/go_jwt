package main

import (
	_ "WechatReport/controllers"
	_ "WechatReport/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Errorhandler("404", controllers.PageNotFound)

	beego.Run()
}
