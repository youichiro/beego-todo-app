package main

import (
	_ "dev/bee-todo-app/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// init DB
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
	}
	err = orm.RegisterDataBase(
		"default",
		"mysql",
		"root:you416570@tcp(127.0.0.1:3306)/beego_todo?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Run()
}

