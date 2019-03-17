package db

//数据库相关

import (
	"time"
	"tzx-admin/lib/log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type callback func()

func MysqlInit(user, pwd, ip, name, runmode string, funcKeepLive callback) {
	e := orm.RegisterDriver("mysql", orm.DRMySQL)
	if e != nil {
		log.Error("Mysql RegisterDriver Err " + e.Error())
		panic(e.Error)
	}
	err := orm.RegisterDataBase("default", "mysql", user+":"+pwd+"@tcp("+ip+")/"+name+"?charset=utf8", 30)
	if err != nil {
		log.Error("Mysql RegisterDataBase Err %s,ip %s", err.Error(), ip)
		panic(err.Error)
	}
	log.Debug("success:%s", user+":"+pwd+"@tcp("+ip+")/"+name+"?charset=utf8")
	//开启自动注册
	orm.RunSyncdb("default", false, true)
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 20)
	db, err := orm.GetDB("default")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	if runmode == "dev" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}
	//开始执行定时检验mysql活性操作
	go func() {
		for {
			funcKeepLive()
			time.Sleep(1 * time.Hour)
		}
	}()
}

func MysqlInitByName(user, pwd, ip, name, runmode string, funcKeepLive callback) {
	e := orm.RegisterDriver("mysql", orm.DRMySQL)
	if e != nil {
		log.Error("Mysql RegisterDriver Err " + e.Error())
		panic(e.Error)
	}
	err := orm.RegisterDataBase(name, "mysql", user+":"+pwd+"@tcp("+ip+")/"+name+"?charset=utf8", 30)
	if err != nil {
		log.Error("Mysql RegisterDataBase Err %s,ip %s", err.Error(), ip)
		panic(err.Error)
	}
	log.Debug("success:%s", user+":"+pwd+"@tcp("+ip+")/"+name+"?charset=utf8")

	orm.SetMaxIdleConns(name, 100)
	if runmode == "dev" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}
	//开始执行定时检验mysql活性操作
	go func() {
		for {
			funcKeepLive()
			time.Sleep(1 * time.Hour)
		}
	}()
}
func MysqlInit2(user, pwd, ip, name, runmode string, funcKeepLive callback) {
	e := orm.RegisterDriver("mysql", orm.DRMySQL)
	if e != nil {
		log.Error("Mysql RegisterDriver Err " + e.Error())
		panic(e.Error)
	}
	err := orm.RegisterDataBase("default", "mysql", user+":"+pwd+"@tcp("+ip+")/"+name+"?charset=utf8", 30)
	if err != nil {
		log.Error("Mysql RegisterDataBase Err %s,ip %s", err.Error(), ip)
		panic(err.Error)
	}
	log.Debug("success:%s", user+":"+pwd+"@tcp("+ip+")/"+name+"?charset=utf8")

	orm.SetMaxIdleConns("default", 100)
	if runmode == "dev" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}
	//开始执行定时检验mysql活性操作
	go func() {
		for {
			funcKeepLive()
			time.Sleep(1 * time.Hour)
		}
	}()
}
