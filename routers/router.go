package routers

import (
	"goLearn/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Get")

	beego.Router("/test", &controllers.ContrController{},"get:Get;post:Post")

	beego.Router("/login", &controllers.LoginController{},"*:Login")
}
