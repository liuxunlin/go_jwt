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

func (u *UserController) URLMapping() {
	u.Mapping("Get", u.Get)
	u.Mapping("Post", u.Post)
	u.Mapping("Redis", u.Redis)
	u.Mapping("Logout", u.Logout)
}

// @router /:uid [get]
func (u *UserController) Get() {
	//u.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", u.Ctx.Request.Header.Get("Origin"))
	fmt.Println("Header Origin: ", u.Ctx.Request.Header.Get("Origin"))
	userId, _ := u.GetInt(":uid")
	if userId == 0 {
		u.Data["json"] = OutResponse(404, nil, "no user exists.")
		u.ServeJSON()
		return
	} else {
		res, err := models.GetUserById(userId)
		if err != nil {
			u.Data["json"] = OutResponse(400, nil, "failed")
			u.ServeJSON()
			return
		}
		//u.Ctx.Output.SetStatus(400)
		u.Data["json"] = OutResponse(200, res, "success")
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
		u.Data["json"] = OutResponse(400, nil, "failed")
		u.ServeJSON()
		return
	}
	u.Data["json"] = OutResponse(200, newId, "success")
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
		u.Data["json"] = OutResponse(200, &user, "success")
		u.ServeJSON()
		return
	} else {
		user, err := models.GetUserById(1)
		if err != nil {
			fmt.Println(err)
		}
		u.Data["json"] = OutResponse(200, &user, "success")
		u.ServeJSON()
		return

	}
	u.Data["json"] = OutResponse(200, &user, "success")
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = OutResponse(200, nil, "logout success.")
	u.ServeJSON()
}
