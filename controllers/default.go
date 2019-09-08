package controllers

import (
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
