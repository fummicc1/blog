package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/blog", &controllers.BlogController{})
}
