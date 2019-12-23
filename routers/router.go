package routers

import (
	"blog/controllers"
	"blog/controllers/blogs"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/blog/new", &blogs.NewBlogController{})
}
