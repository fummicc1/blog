package controllers

import (
	"blog/models"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// BlogController handles "/blog" routing.
type BlogController struct {
	beego.Controller
}

var blogManager *models.BlogManager

func init() {
	blogManager = &models.BlogManager{
		O: orm.NewOrm(),
	}
}

// Get is handler for "GET /blog".
func (controller *BlogController) Get() {
	logs.Debug("GET /blog")
	controller.Data["blogs"] = blogManager.Read()
	controller.TplName = "blog.html"
}

func (controller *BlogController) Post() {
	logs.Debug("Post /blog")
	var blog models.Blog = models.Blog{
		Date: time.Now(),
	}
	err := controller.ParseForm(&blog)
	if err != nil {
		logs.Error("failed to parse form: &s", err)
		return
	}
	err = blogManager.Create(&blog)
	if err != nil {
		logs.Error("failed to create blog to mysql: %s", err)
		return
	}
	controller.Data["blogs"] = blogManager.Read()
	controller.TplName = "blog.html"
}
