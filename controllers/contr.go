package controllers

import "github.com/astaxie/beego"

type ContrController struct {
	beego.Controller
}

type UserInfo struct {
	Username string
	Password string
}

func (c *ContrController) Get() {
	//c.Ctx.WriteString("<p style=\"color:red\">hello world</p>")
	//id := c.GetString("id")
	//c.Ctx.WriteString(id)
	//name := c.GetString("name")
	//c.Ctx.WriteString("<br/>")
	//c.Ctx.WriteString(name)


	//name := c.Ctx.GetCookie("name")
	//password := c.Ctx.GetCookie("password")

	name := c.GetSession("name")
	password := c.GetSession("password")

	if nameString,ok := name.(string);ok &&nameString != ""{
		c.Ctx.WriteString("name is:"+nameString+"   password is :"+password.(string))
	}else {
		c.Ctx.WriteString(`<html><form action="http://localhost:8080/test" method ="post">
					<input type="text" name="Username"/>
					<input type="password" name="Password"/>
					<input type="submit" value = "提交"/>
				</form></html>`)
	}
}

func (c *ContrController) Post(){

	user := UserInfo{}
	err := c.ParseForm(&user)
	if err != nil{
		panic(err)
	}


	c.Ctx.SetCookie("name",user.Username,100,"/")
	c.Ctx.SetCookie("password",user.Password,100,"/")
	c.SetSession("name",user.Username)
	c.SetSession("password",user.Password)

	c.Ctx.WriteString("name is:"+user.Username+"   password is :"+user.Password)
}