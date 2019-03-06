package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Main

type MainController struct {
	beego.Controller
}

//map[string]interface{}{"code": 400, "msg": "no user exists!", "data": nil}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (m *MainController) Welcome() {
	m.Data["Website"] = "www.unclepang.com"
	m.Data["Email"] = "10846295@qq.com"
	m.TplName = "index.tpl"
	m.Render()
}
