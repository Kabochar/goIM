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
	drivername := "mysql"
	Dsname := "root:mysql123@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DBEngin, err = xorm.NewEngine(drivername, Dsname)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	// 是否显示 SQL 语句
	DBEngin.ShowSQL(true)
	// 数据库最大链接数
	DBEngin.SetMaxOpenConns(2)

	// 自动user
	err = DBEngin.Sync2(new(model.User))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("init DataBase OK...")
}
