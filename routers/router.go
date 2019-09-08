package routers

import (
	"github.com/astaxie/beego"
	"iview-admin/controllers"
)

func init() {
	beego.AutoRouter(&controllers.Controller{})
	beego.AutoRouter(&controllers.LoginController{})
	//beego.Router("/", &controllers.MainController{})
}
