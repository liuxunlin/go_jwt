// @APIVersion 1.0.0
// @Title go jwt api
// @Description go jwt api documents
// @Contact 10846295@qq.com
// @TermsOfServiceUrl http://www.unclepang.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go_jwt/controllers"
	"go_jwt/middlewares"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	// corsHandler := func(ctx *context.Context) {
	// 	ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Domain())
	// 	ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	// 	fmt.Println("allow origin =>>> ", ctx.Input.Domain())
	// }
	// beego.InsertFilter("*", beego.BeforeRouter, corsHandler)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Authorization,Accept,Accept-Encoding, Authorization, Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{}, "*:Welcome")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSBefore(middlewares.ValidateToken),
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
