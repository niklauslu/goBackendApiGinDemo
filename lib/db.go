package lib

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// DbEngine ...
var engine *xorm.Engine

// DBConnect 连接数据库
func DBConnect(driverName string, dataSourceName string, debug string) (err error) {
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		// log.Fatalf("db connect err %s", err.Error())
		return err
	}

	if debug == "true" {
		engine.ShowSQL(true)
	}

	return nil
}

// DBSessionGet 获取连接会话
func DBSessionGet() (session *xorm.Session) {
	session = engine.NewSession()
	return session
}

// DBSync 数据库同步
func DBSync(models ...interface{}) error {
	err := engine.Sync(models...)
	return err
}

// DBStartTransation 开启事务
func DBStartTransation(session *xorm.Session) (*xorm.Session, error) {
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return session, err
	}
	return session, nil
}
