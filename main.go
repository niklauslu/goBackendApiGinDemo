package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var logger = lib.Logger("app")
var uploadPath = ""

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	setLogger()

	setDatabase()

	router := gin.Default()

	// 设置上传
	if err := setUploader(router); err != nil {
		logger.Fatalf("upload set err:%s", err.Error())
	}
	// 全局中间件
	router.Use(globalConfMid())
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	router.Use(gin.Recovery())

	setRouter(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Println("Server exiting")
}

func setLogger() {
	// 记录到文件。
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	gin.ForceConsoleColor()
}

func setDatabase() {
	err := lib.DBConnect("mysql", os.Getenv("DB_DSN"), os.Getenv("DEBUG"))
	if err != nil {
		logger.Fatalf("db connnet err: %s", err.Error())
	}
	logger.Info("db connnet success")

	err = lib.DBSync(
		new(model.TUser),
	)

	if err != nil {
		logger.Errorf("db sync err: %s", err.Error())
	}
	logger.Info("db sync success")

}

func setUploader(router *gin.Engine) error {

	if dir, err := os.Getwd(); err == nil {
		uploadPath = dir + "/public/uploads/"
	}

	if err := os.MkdirAll(uploadPath, 0777); err != nil {
		return err
	}

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/uploads/", uploadPath)

	return nil
}

// 全局配置中间件
func globalConfMid() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set("upload_path", uploadPath)
		c.Set("upload_host", os.Getenv("UPLOAD_HOST"))

		//请求之前
		c.Next()
	}
}
