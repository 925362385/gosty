package controllers

import (
	"github.com/astaxie/beego"
	"goLearn/models"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

func (login *LoginController) Login() {

	//login.Ctx.WriteString("login")
	user := models.User{Username: "ww", Password: "qwe"}
	i, e := models.AddUser(&user)
	if e != nil {
		panic(e)
	}
	login.Ctx.WriteString(strconv.Itoa(int(i)))
}
