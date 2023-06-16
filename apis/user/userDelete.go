package apis_user

import (
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"

	"github.com/gin-gonic/gin"
)

func UserDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	sess := lib.DBSessionGet()
	defer sess.Close()

	var user model.TUser

	sess.ID(id).Get(&user)
	if user.ID <= 0 {
		ctx.JSON(400, gin.H{"message": "数据错误"})
		return
	}

	user.Status = -1
	_, err := sess.ID(1).Cols("status").Update(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}
