package main

import (
	"fmt"
	"net/http"
	apis_article "niklauslu/goBackendApiGinDemo/apis/article"
	apis_upload "niklauslu/goBackendApiGinDemo/apis/upload"
	apis_user "niklauslu/goBackendApiGinDemo/apis/user"
	"time"

	"github.com/gin-gonic/gin"
)

func setRouter(router *gin.Engine) {
	getApiRouter(router)
	getAdminRouter(router)

}

func getApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/timestamp", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, fmt.Sprintf("%d", time.Now().Unix()))
		})

		api.POST("/upload/local", apis_upload.SingleLocalUpload)

		api.GET("/users", apis_user.UsersGet)
		api.GET("/users/:id", apis_user.UserGet)
		api.POST("/users", apis_user.UserCreate)
		api.PUT("/users", apis_user.UserUpdate)
		api.DELETE("/users/:id", apis_user.UserDelete)

		api.POST("/articles", apis_article.ArticleAdd)
		api.GET("/articles", apis_article.ArticleGet)
	}
}

func getAdminRouter(router *gin.Engine) {
	admin := router.Group("/api/admin")
	{
		admin.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}
}
