package controllers

import (
	"github.com/astaxie/beego"
)

// CommonController operations for Common
type CommonController struct {
	beego.Controller
}

// URLMapping ...
func (c *CommonController) URLMapping() {
	c.Mapping("PageNotFound", c.PageNotFound)
	c.Mapping("GetAll", c.GetAll)
}

// PageNotFound ...
// @Title GetOne
// @Description get Common by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Common
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommonController) PageNotFound() {

	c.Data["json"] = "Page Not Found"
	c.Data["code"] = 404
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Common
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Common
// @Failure 403
// @router / [get]
func (c *CommonController) GetAll() {

}
