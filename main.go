package main

import (
	"fmt"
	"go_wechat/controllers"
	"go_wechat/libs"
	_ "go_wechat/routers"
	"runtime"
	"time"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	libs.Init()
	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	//指定使用多核，核心数为CPU的实际核心数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	// 记录启动时间
	beego.AppConfig.Set("up_time", fmt.Sprintf("%d", time.Now().Unix()))
	beego.ErrorController(&controllers.ErrorController{})
	orm.DefaultTimeLoc = time.UTC

	beego.Run()
}
