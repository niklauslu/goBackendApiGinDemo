 # golang接口服务demo（gin框架版本）

项目使用：  
+ [gin](https://github.com/gin-gonic/gin):一个非常好用的go web框架  
+ [xorm](https://xorm.io/):golang的ORM，操作数据库很好用
+ [docker](https://www.docker.com/):打包部署
+ [ACME](https://github.com/acmesh-official/acme.sh):获取https证书用

安装gin
```sh
go get -u github.com/gin-gonic/gin
```
ps:设置go env
```sh
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

gin demo示例
```go
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

### 开发

#### [项目环境配置（.env）](./docs/env.md)
#### [日志文件（logger）](./docs/log.md)
#### [全局配置中间件](./docs/globalConf.md)
#### [设置路由（router）](./docs/router.md)
#### [数据库使用（database)](./docs/database.md)
#### [文件上传（upload）]()

### 数据库使用

#### [mysql](./docs/mysql.md)
#### [MongoDB](./docs/mongodb.md)
### restful api 示例

[具体示例看这里](./docs/restful.md)




### 部署

部署基于[docker](https://www.docker.com/)，docker安装教程可见[官方文档](https://docs.docker.com/get-docker/)

1. [打包docker镜像](./docs/build.md)
2. [服务器部署](./docs/deploy.md)
3. [nginx&域名相关](./docs/domain.md)
