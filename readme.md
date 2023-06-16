 # golang接口服务demo（gin框架版本）

[gin](https://github.com/gin-gonic/gin):一个非常好用的go web框架

安装gin
```
go get -u github.com/gin-gonic/gin
```
ps:设置go env
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

gin demo示例
```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

### 开发准备

#### [项目环境配置（.env）](./docs/env.md)
#### [日志文件（logger）](./docs/log.md)
#### [设置路由（router）](./docs/router.md)
#### [数据库使用（database)](./docs/database.md)

#### restful api 示例

[具体示例看这里](./docs/restful.md)

#### 部署