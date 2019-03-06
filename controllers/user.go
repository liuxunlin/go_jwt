package controllers

import (
	"encoding/json"
	"fmt"
	"go_wechat/libs"
	"go_wechat/models"
	"time"

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
		u.Data["json"] = &Response{404, "no user exists.", nil}
		u.ServeJSON()
		return
	} else {
		res, err := models.GetUserById(userId)
		if err != nil {
			u.Data["json"] = &Response{400, "failed", nil}
			u.ServeJSON()
			return
		}
		u.Data["json"] = &Response{200, "success", &res}
		u.ServeJSON()
		return
	}
}

// @router /create [post]
func (u *UserController) Post() {
	var NewUser models.User

	NewUser.Name = u.GetString("Name")
	NewUser.Age, _ = u.GetInt("Age")
	NewUser.Email = u.GetString("Email")
	NewUser.CreatedAt = time.Now().Unix()
	NewUser.UpdatedAt = time.Now().Unix()
	newId, err := models.AddUser(&NewUser)

	if err != nil {
		u.Data["json"] = &Response{400, "failed", nil}
		u.ServeJSON()
		return
	}
	u.Data["json"] = &Response{200, "success", newId}
	u.ServeJSON()
	return
}

// @router /redis/:redisKey [get]
func (u *UserController) Redis() {
	var user models.User
	redisKey := u.GetString(":redisKey")
	userInfo, err := libs.GetKey(redisKey)
	json.Unmarshal([]byte(userInfo), &user)
	if err != nil {
		user, err := models.GetUserById(1)
		userJson, _ := json.Marshal(&user)
		libs.SetKey(redisKey, userJson, 3600)
		if err != nil {
			fmt.Println(err)
		}
		u.Data["json"] = &Response{200, "success", &user}
		u.ServeJSON()
		return
	} else {
		user, err := models.GetUserById(1)
		if err != nil {
			fmt.Println(err)
		}
		u.Data["json"] = &Response{200, "success", &user}
		u.ServeJSON()
		return

	}
	u.Data["json"] = &Response{200, "success", &user}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = &Response{200, "logout success.", nil}
	u.ServeJSON()
}
