package v1

import (
	"test/pkg/app"
	"test/pkg/enum"

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

}

func Register(c *gin.Context) {
	// appG := app.Gin{C: c}

}
