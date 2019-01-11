package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/golangDemo/webApiDemo/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/golangDemo/webApiDemo/controllers:PayController"],
		beego.ControllerComments{
			Method: "PayAdd",
			Router: `/PayAdd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/golangDemo/webApiDemo/controllers:PayController"] = append(beego.GlobalControllerRouter["github.com/golangDemo/webApiDemo/controllers:PayController"],
		beego.ControllerComments{
			Method: "PayQuery",
			Router: `/PayQuery`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
