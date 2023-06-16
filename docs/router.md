### 路由设置

```go
// file: router.go
// 封装路由方法

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setRouter(router *gin.Engine) {

    api.GET("/timestamp", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("%d", time.Now().Unix()))
	})

}
```

```go
// file: main.go
router := gin.Default()
// 调用setRouter方法
setRouter(router)
```

#### 设置路由分组

```go
// file: router.go

func setRouter(router *gin.Engine) {
	getApiRouter(router)
}

func getApiRouter(router *gin.Engine) {
    // 路由api分组
	api := router.Group("/api")
	{
		api.GET("/timestamp", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, fmt.Sprintf("%d", time.Now().Unix()))
		})
	}
}
```

