package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
var db orm.Ormer

type User struct {
	Id       int
	Username string
	Password string
}

func init() {
	debug := beego.AppConfig.DefaultBool("Debug", false)
	sqlconn := beego.AppConfig.DefaultString("sqlconn", "root:9253@tcp(localhost:3306)/graprod?charset=utf8")
	orm.Debug = debug
	orm.RegisterDataBase("default", "mysql", sqlconn)
	orm.RegisterModel(new(User))
	db = orm.NewOrm()
}

func AddUser(user *User) (int64, error) {
	i, e := db.Insert(user)

	return i, e
}
