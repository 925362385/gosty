package main

import (
	"github.com/astaxie/beego"
	_ "goLearn/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
