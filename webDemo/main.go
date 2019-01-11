package main

import (
	"github.com/golangDemo/webDemo/models"

	"github.com/astaxie/beego"
	_ "github.com/golangDemo/webDemo/routers"
)

func main() {
	defer models.Close()

	beego.Run()

}
