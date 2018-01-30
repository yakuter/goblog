package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/goblog?charset=utf8")
	orm.RunSyncdb("default", false, true)
}
