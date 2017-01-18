// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"driiip-api/controllers/device"
	"driiip-api/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/device",
			beego.NSInclude(
				&device.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
