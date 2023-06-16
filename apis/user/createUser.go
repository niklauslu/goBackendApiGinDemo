package apis_user

import (
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/model"
	"niklauslu/goBackendApiGinDemo/utils"

	"github.com/gin-gonic/gin"
)

type createUserBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   int32  `json:"status"`
}

func UserCreate(ctx *gin.Context) {
	var body createUserBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	sess := lib.DBSessionGet()
	defer sess.Close()

	var user model.TUser
	user.Username = body.UserName
	user.Email = body.Email
	user.Password = utils.MD5(body.Password)
	user.Role = body.Role
	user.Status = body.Status

	_, err := sess.Insert(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}
