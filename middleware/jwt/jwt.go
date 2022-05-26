package jwt

import (
	"net/http"
	"test/lib"
	"test/pkg/enum"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = enum.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = enum.INVALID_PARAMS
		} else {
			_, err := lib.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = enum.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = enum.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != enum.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  enum.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
