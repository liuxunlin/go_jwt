package controllers

import "github.com/astaxie/beego"

// ErrorController definition.
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = OutResponse(404, nil, "Not Found")
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = OutResponse(401, nil, "Permission denied")
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = OutResponse(403, nil, "Forbidden")
	c.ServeJSON()
}
