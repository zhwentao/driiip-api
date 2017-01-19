package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpMonitorLogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpPlantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"] = append(beego.GlobalControllerRouter["driiip-api/controllers:DpUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
