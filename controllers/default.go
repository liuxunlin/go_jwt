package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type MainController struct {
	beego.Controller
}

// Get method.
func (c *MainController) Get() {

	data := map[string]string{
		"code": "200",
		"info": "success",
		"data": "",
	}
	c.Data["json"] = data
	c.ServeJSON()
}
