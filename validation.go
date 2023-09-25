package gofusion

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func getErrorMessage(e validator.FieldError) string {
	message := "undefined error"

	if e.Tag() == "required" {
		message = "required"
	}

	return message
}

func Validate(c *gin.Context, request any) map[string]string {
	if err := c.ShouldBindJSON(request); err != nil {
		v := validator.New()

		err := v.Struct(request)
		errors := make(map[string]string, len(err.(validator.ValidationErrors)))
		for _, e := range err.(validator.ValidationErrors) {
			errors[toSnakeCase(e.Field())] = getErrorMessage(e)
		}

		return errors
	}

	return nil
}

// middleware
func ValidateURIParam(paramName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramValue := ctx.Param(paramName)
		if paramValue == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": paramName + "_is_required",
				"success": false,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
