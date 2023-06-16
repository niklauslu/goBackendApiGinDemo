package apis_user

import (
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"
	"niklauslu/goBackendApiGinDemo/utils"

	"github.com/gin-gonic/gin"
)

type updateUserBody struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   int32  `json:"status"`
}

func UserUpdate(ctx *gin.Context) {
	var body updateUserBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	id := body.ID
	if id <= 0 {
		ctx.JSON(400, gin.H{"message": "参数错误"})
		return
	}

	sess := lib.DBSessionGet()
	defer sess.Close()

	var user model.TUser

	sess.ID(id).Get(&user)
	if user.ID <= 0 {
		ctx.JSON(400, gin.H{"message": "数据错误"})
		return
	}

	user.Username = body.UserName
	user.Email = body.Email
	// user.Password = utils.MD5(body.Password)
	user.Role = body.Role
	user.Status = body.Status

	if body.Password != "" {
		user.Password = utils.MD5(body.Password)
	}

	_, err := sess.ID(1).AllCols().Update(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}
