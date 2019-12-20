package controllers

import (
	"github.com/astaxie/beego"
)

// MainController handles "/" routing.
type MainController struct {
	beego.Controller
}

// Get is handler for "GET /".
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
