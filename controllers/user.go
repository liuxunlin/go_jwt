package controllers

import (
	"go_wechat/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @router /:uid [get]
func (u *UserController) Get() {
	userId, _ := u.GetInt(":uid")
	if userId == 0 {
		u.Data["json"] = "no user exists."
		u.ServeJSON()
		return
	} else {
		res, err := models.GetUserById(userId)
		if err != nil {
			u.Data["json"] = err
			u.ServeJSON()
			return
		}
		u.Data["json"] = res
		u.ServeJSON()
		return
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
