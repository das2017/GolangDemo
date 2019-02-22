package main

import (
	"github.com/GolangDemo/webDemo/models"

	"github.com/astaxie/beego"
	_ "github.com/GolangDemo/webDemo/routers"
)

func main() {
	defer models.Close()

	beego.Run()

}
