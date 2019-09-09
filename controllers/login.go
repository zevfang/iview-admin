package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}
