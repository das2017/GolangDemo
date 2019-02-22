package routers

import (
	"github.com/astaxie/beego"
	"github.com/GolangDemo/webDemo/controllers"
)

func init() {
	beego.Router("/query", &controllers.MainController{}, "Post:PayQuery")
	beego.Router("/add", &controllers.MainController{}, "Post:PayAdd")
	beego.Router("/", &controllers.MainController{}, "Get:Home")
}
