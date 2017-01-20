// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"driiip-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/dp_user_device",
			beego.NSInclude(
				&controllers.DpUserDeviceController{},
			),
		),
		beego.NSNamespace("/dp_user",
			beego.NSInclude(
				&controllers.DpUserController{},
			),
		),
		beego.NSNamespace("/dp_plant",
			beego.NSInclude(
				&controllers.DpPlantController{},
			),
		),
		beego.NSNamespace("/dp_monitor_log",
			beego.NSInclude(
				&controllers.DpMonitorLogController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
