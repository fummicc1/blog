package controllers

import (
	"blog/models"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

var logger *logs.BeeLogger = beego.BeeLogger

var blogManager = models.BlogManager{
	O: orm.NewOrm(),
}

// BlogController handles "/blog" routing.
type BlogController struct {
	beego.Controller
}

// Get is handler for "GET /blog".
func (controller *BlogController) Get() {
	logger.Debug("GET /blog")
	controller.Data["blogs"] = blogManager.Read()
	controller.TplName = "blog.html"
}

// Post is handler for "POST /blog".
func (controller *BlogController) Post() {

}
