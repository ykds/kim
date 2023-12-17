package middleware

import (
	"github.com/gin-gonic/gin"
	"kim/internal/logic/errcode"
	"kim/internal/logic/global"
	"kim/internal/pkg/jwt"
	"kim/internal/pkg/response"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader(global.Token)
		if token == "" {
			response.HandleResponse(context, errcode.InvalidTokenErr, nil)
			context.Abort()
			return
		}
		userId, err := jwt.ParseToken(token)
		if err != nil {
			response.HandleResponse(context, errcode.InvalidTokenErr, nil)
			context.Abort()
			return
		}
		context.Set(global.UserIdKey, userId)
		context.Next()
	}
}
