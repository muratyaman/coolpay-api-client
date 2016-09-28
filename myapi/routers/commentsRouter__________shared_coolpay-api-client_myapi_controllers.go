package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["myapi/controllers:CoolpayController"] = append(beego.GlobalControllerRouter["myapi/controllers:CoolpayController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CoolpayController"] = append(beego.GlobalControllerRouter["myapi/controllers:CoolpayController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CoolpayController"] = append(beego.GlobalControllerRouter["myapi/controllers:CoolpayController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CoolpayController"] = append(beego.GlobalControllerRouter["myapi/controllers:CoolpayController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:CoolpayController"] = append(beego.GlobalControllerRouter["myapi/controllers:CoolpayController"],
		beego.ControllerComments{
			Method: "RequestLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:UserController"] = append(beego.GlobalControllerRouter["myapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
