### 全局配置中间件

全局配置中间件函数
```go
// file: main.go
func globalConfMid() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set("KEY", "VALUE")

		//请求之前
		c.Next()
	}
}
```


```go
// file:main.go 加载中间件
router.Use(globalConfMid())
```

在接口函数中调用示例
```go
func APIDemo(c *gin.Context) {
    value := c.GetString("KEY")
}
```

