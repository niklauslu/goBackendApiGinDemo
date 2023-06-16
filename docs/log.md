### 日志

#### 设置gin框架日志

```go
// file: main.go

func setLogger() {
	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	gin.ForceConsoleColor()
}

```

#### 日志文件

[logrus](https://github.com/sirupsen/logrus)

```
go get github.com/sirupsen/logrus
```

封装logger方法
```go
// file： lib/logger.go
package lib

import (
	"os"
	"path"
	"github.com/sirupsen/logrus"
	"fmt"
  "time"
  
  rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// Logger 日志
func Logger(name string) *logrus.Logger {
  // now := time.Now()
  logFilePath := ""
  if dir, err := os.Getwd(); err == nil {
    logFilePath = dir + "/logs/"
  }
  if err := os.MkdirAll(logFilePath, 0777); err != nil {
    fmt.Println(err.Error())
  }
  // logFileName := name + "_" + now.Format("2006-01-02") + ".log"
  logFileName := name + ".log"
  //日志文件
  fileName := path.Join(logFilePath, logFileName)
  if _, err := os.Stat(fileName); err != nil {
    if _, err := os.Create(fileName); err != nil {
      fmt.Println(err.Error())
    }
  }
  //写入文件
  // src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  // if err != nil {
  //   fmt.Println("err", err)
  // }

  writer, _ := rotatelogs.New(
    fileName+".%Y%m%d%H%M",
    rotatelogs.WithLinkName(fileName),
    rotatelogs.WithMaxAge(7*24*time.Hour),
    rotatelogs.WithRotationTime(24*time.Hour),
    )

  //实例化
  logger := logrus.New()

  //设置输出
  // logger.Out = src
  logger.SetOutput(writer)

  //设置日志级别
  logger.SetLevel(logrus.DebugLevel)


  //设置日志格式
  logger.SetFormatter(&logrus.TextFormatter{
    TimestampFormat: "2006-01-02 15:04:05",
  })
  
  return logger
}
```
