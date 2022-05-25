package app

import (
	"net/http"
	"test/pkg/enum"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, enum.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, enum.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, enum.INVALID_PARAMS
	}

	return http.StatusOK, enum.SUCCESS
}
