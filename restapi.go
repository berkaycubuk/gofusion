package gofusion

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondJSON(ctx *gin.Context, status int, message string, data interface{}) {
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	ctx.JSON(status, response)
}

// middleware
func HandleAPIResponse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			var errors []string
			for _, err := range ctx.Errors {
				errors = append(errors, err.Error())
			}
			RespondJSON(ctx, http.StatusInternalServerError, "error", errors)
			ctx.Abort()
			return
		}

		if ctx.Writer.Status() == http.StatusOK {
			data, exists := ctx.Get("response")
			if exists {
				RespondJSON(ctx, http.StatusOK, "success", data)
				ctx.Abort()
				return
			}
		}
	}
}
