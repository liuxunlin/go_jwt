package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go_wechat/controllers:UserController"] = append(beego.GlobalControllerRouter["go_wechat/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go_wechat/controllers:UserController"] = append(beego.GlobalControllerRouter["go_wechat/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go_wechat/controllers:UserController"] = append(beego.GlobalControllerRouter["go_wechat/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go_wechat/controllers:UserController"] = append(beego.GlobalControllerRouter["go_wechat/controllers:UserController"],
        beego.ControllerComments{
            Method: "Redis",
            Router: `/redis/:redisKey`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
