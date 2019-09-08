package main

import (
	"github.com/astaxie/beego"
	_ "iview-admin/routers"
)

func main() {
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.Run()
}
