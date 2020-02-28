package main

import (
	_ "achievement/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/beego.log"}`)
	orm.RegisterDataBase("default","mysql","root:123456@tcp(192.168.0.141:3306)/ssms?charset=utf8")
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run();
}

