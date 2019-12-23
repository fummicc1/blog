package blogs

import (
	"blog/models"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var blogManager *models.BlogManager

type NewBlogController struct {
	beego.Controller
}

func init() {
	blogManager = &models.BlogManager{
		O: orm.NewOrm(),
	}
}

func (c *NewBlogController) Get() {
	logs.Debug("Get /blog/new")
	c.TplName = "blog/create.html"
}
