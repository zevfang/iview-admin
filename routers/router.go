package routers

import (
	"github.com/astaxie/beego"
	"iview-admin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
