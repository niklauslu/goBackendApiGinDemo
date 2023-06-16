package apis_user

import (
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"

	"github.com/gin-gonic/gin"
)

var logger = lib.Logger("user")

// type userGetParams struct {
// 	ID string `uri:"id" binding:"required"`
// }

func UserGet(ctx *gin.Context) {
	// var params userGetParams

	// if err := ctx.ShouldBindUri(&params); err != nil {
	// 	ctx.JSON(400, gin.H{"message": err.Error()})
	// 	return
	// }

	id := ctx.Param("id")
	logger.Infof("/api/user/%s", id)

	sess := lib.DBSessionGet()
	defer sess.Close()

	var user model.TUser

	sess.ID(id).Get(&user)

	ctx.JSON(http.StatusOK, user)
}
