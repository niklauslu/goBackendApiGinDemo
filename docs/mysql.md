### 数据库使用(mysql)

以[xorm](https://xorm.io/)（A Simple and Powerful ORM for Go）为例，mysql版本

#### 安装
```
go get xorm.io/xorm
go get github.com/go-sql-driver/mysql
```

#### 封装db方法

```go
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
```

主函数设置
```go
func setDatabase() {
	err := lib.DBConnect("mysql", os.Getenv("DB_DSN"), os.Getenv("DEBUG"))
	if err != nil {
		logger.Fatalf("db connnet err: %s", err.Error())
	}
	logger.Info("db connnet success")

    // 同步数据库
    err = lib.DBSync(
		new(...),
	)

	if err != nil {
		logger.Errorf("db sync err: %s", err.Error())
	}
	logger.Info("db sync success")
}

```