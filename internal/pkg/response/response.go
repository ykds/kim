package response

import (
	"github.com/gin-gonic/gin"
	"kim/internal/pkg/errors"
	"kim/internal/pkg/log"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		var e = new(errors.Error)
		if errors.As(err, e) {
			ctx.JSON(200, Response{Code: e.Code(), Message: e.Message()})
		} else {
			ctx.JSON(200, Response{Code: errors.InternalError.Code(), Message: errors.InternalError.Message()})
		}
		return
	}
	ctx.JSON(200, Response{Code: errors.Success.Code(), Message: errors.Success.Message(), Data: data})
}

func LogResponse(log *log.Logger, err error) {
	if err != nil {
		log.Errorf("%+v\n", err)
	}
}
