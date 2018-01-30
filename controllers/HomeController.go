package controllers

import (
"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Welcome"] = "ahan da oldu"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}