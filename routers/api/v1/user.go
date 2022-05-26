package v1

import (
	"test/lib"
	"test/models"
	"test/pkg/app"
	"test/pkg/enum"
	userservice "test/service/user_service"

	"github.com/gin-gonic/gin"
)

type UserJSON struct {
	Username string `json:"username" valid:"Required;MaxSize(255)"`
	Passowrd string `json:"password" valid:"Required;MaxSize(255)"`
}

func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		user UserJSON
	)

	httpCode, errCode := app.BindAndValid(c, &user)
	if errCode != enum.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	mUser := &userservice.User{
		Username: user.Username,
		Password: user.Passowrd,
	}

	userData, err := mUser.GetUserByUsername()

	if err != nil || (userData == models.User{}) {
		appG.Response(enum.SUCCESS, enum.LOGIN_FAIL, err)
		return
	}

	token, err := lib.GenerateToken(mUser.Username, mUser.Password)

	if userData.Password != lib.MD5(mUser.Password) {
		appG.Response(enum.SUCCESS, enum.LOGIN_FAIL, nil)
		return
	}

	appG.Response(enum.SUCCESS, enum.SUCCESS, map[string]interface{}{
		"token": token,
	})
}

func Register(c *gin.Context) {
	// appG := app.Gin{C: c}

}
