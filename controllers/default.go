package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Users struct {
	Username string
	Password string
}


func (c *MainController) Get() {
	user := Users{}

	if err := c.ParseForm(&user);err!= nil{

	}

	c.Ctx.WriteString("Username:"+user.Username+"Password:"+user.Password)
}
