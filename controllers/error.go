package controllers

import "github.com/astaxie/beego"

// ErrorController definition.
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = Response{
		Code: 404,
		Msg:  "Not Found",
		Data: nil,
	}
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = Response{
		Code: 401,
		Msg:  "Permission denied",
		Data: nil,
	}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = Response{
		Code: 403,
		Msg:  "Forbidden",
		Data: nil,
	}
	c.ServeJSON()
}
