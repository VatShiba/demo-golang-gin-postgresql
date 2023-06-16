package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Validate(s interface{}, ctx *gin.Context) (err any) {
	v := validator.New()
	if err := v.Struct(s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return err.Error()
	}
	return nil
}
