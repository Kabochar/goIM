package service

import (
	"errors"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/go-xorm/xorm"
	"goIM/model"
	"log"
)

var DBEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:mysql123@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DBEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DBEngin.ShowSQL(true)
	//数据库最大打开的连接数B
	DBEngin.SetMaxOpenConns(2)

	//自动User
	DBEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}
