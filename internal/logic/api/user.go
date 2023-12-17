package api

import (
	"github.com/gin-gonic/gin"
	"kim/internal/logic/global"
	"kim/internal/logic/middleware"
	"kim/internal/logic/service"
	"kim/internal/pkg/errors"
	"kim/internal/pkg/response"
)

func Register(ctx *gin.Context) {
	var (
		req  service.RegisterReq
		resp service.RegisterResp
		err  error
	)
	defer func() {
		response.LogResponse(global.Logger, err)
		response.HandleResponse(ctx, err, resp)
	}()
	if err = ctx.BindJSON(&req); err != nil {
		err = errors.BadParameters
		return
	}
	resp, err = services.UserService.Register(req)
}

func Login(ctx *gin.Context) {
	var (
		req  service.LoginReq
		resp service.LoginResp
		err  error
	)
	defer func() {
		response.LogResponse(global.Logger, err)
		response.HandleResponse(ctx, err, resp)
	}()
	if err = ctx.BindJSON(&req); err != nil {
		err = errors.BadParameters
		return
	}
	resp, err = services.UserService.Login(req)
}

func Logout(ctx *gin.Context) {
	var (
		req  service.LogoutReq
		resp service.LogoutResp
		err  error
	)
	defer func() {
		response.LogResponse(global.Logger, err)
		response.HandleResponse(ctx, err, resp)
	}()
	userId, _ := ctx.Get(global.UserIdKey)
	req.UserId = userId.(uint)
	resp, err = services.UserService.Logout(req)
}

func registerUserRouter(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/register", Register)
		user.POST("/login", Login)
		user.POST("/logout", middleware.Auth(), Logout)
	}
}
