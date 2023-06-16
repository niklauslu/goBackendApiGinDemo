package apis_user

import (
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type UsersGetQuery struct {
// 	Page   string `form:"page"`
// 	Size   string `form:"size"`
// 	Search string `form:"search"`
// }

func UsersGet(ctx *gin.Context) {
	// var query UsersGetQuery
	// logger.Infof("/api/users , %+v", query)
	// if err := ctx.ShouldBindQuery(&query); err != nil {
	// 	logger.Errorf("/api/users err: %s", err.Error())
	// 	ctx.JSON(400, gin.H{"message": err.Error()})
	// }

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	search := ctx.Query("search")

	logger.Infof("/api/users page:%d, size:%d, search: %s", page, size, search)

	sess := lib.DBSessionGet()
	defer sess.Close()

	sess = sess.Where("status >= 0")

	if search != "" {
		sess = sess.Where("username like %?% or email like %?%", search, search)
	}

	if size > 0 {
		offset := (page - 1) * size
		sess = sess.Limit(page, offset)
	}

	var users []model.TUser
	count, _ := sess.FindAndCount(&users)

	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
		"rows":  users,
	})
}
