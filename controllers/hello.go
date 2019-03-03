package controllers

import (
	"github.com/astaxie/beego"
)

// HelloController operations for Hello
type HelloController struct {
	beego.Controller
}

// @Title Get
// @Description get Hello
// @Success 200 {object} models.Hello
// @Failure 403
// @router /hello [get]
func (c *HelloController) Get() {
	c.Data["json"] = "Hello hello!!!"
	c.ServeJSON()
}
