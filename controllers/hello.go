package controllers

import (
	"github.com/astaxie/beego"
)

func (c *HelloController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// HelloController operations for Hello
type HelloController struct {
	beego.Controller
}

// @Title Get
// @Description get Hello
// @Success 200 {object} models.Hello
// @Failure 403
// @router /:name [get]
func (c *HelloController) Get() {
	name := c.GetString(":name")
	if name != "" {
		c.Data["json"] = "Hello，" + name
	} else {
		c.Data["json"] = "Hello Gender"
	}
	c.ServeJSON()
}
