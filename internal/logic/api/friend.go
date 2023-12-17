package api

import (
	"github.com/gin-gonic/gin"
	"kim/internal/logic/global"
	"kim/internal/logic/middleware"
	"kim/internal/logic/service"
	"kim/internal/pkg/errors"
	"kim/internal/pkg/response"
)

func ApplyFriend(ctx *gin.Context) {
	var (
		req  service.FriendApplyReq
		resp service.FriendApplyResp
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
	userId, _ := ctx.Get(global.UserIdKey)
	req.UserId = userId.(uint)
	resp, err = services.FriendApplicationService.FriendApply(req)
}

func HandleApplyStatus(ctx *gin.Context) {
	var (
		req  service.UpdateApplyReq
		resp service.UpdateApplyResp
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
	userId, _ := ctx.Get(global.UserIdKey)
	resp, err = services.FriendApplicationService.UpdateFriendApplicationStatus(userId.(uint), req)
}

func ApplyQuestion(ctx *gin.Context) {
	var (
		req  service.ApplicationQuestReq
		resp service.ApplicationQuestResp
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
	userId, _ := ctx.Get(global.UserIdKey)
	resp, err = services.FriendApplicationService.ApplicationQuest(userId.(uint), req)
}

func ApplyAnswer(ctx *gin.Context) {
	var (
		req  service.ApplicationAnsReq
		resp service.ApplicationAnsResp
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
	userId, _ := ctx.Get(global.UserIdKey)
	resp, err = services.FriendApplicationService.ApplicationAns(userId.(uint), req)
}

func ListApplication(ctx *gin.Context) {
	var (
		resp struct {
			List []*service.ApplicationInfo `json:"list"`
		}
		err error
	)
	defer func() {
		response.LogResponse(global.Logger, err)
		response.HandleResponse(ctx, err, resp)
	}()
	userId, _ := ctx.Get(global.UserIdKey)
	resp.List, err = services.FriendApplicationService.ListApplication(userId.(uint))
}

func ListFriend(ctx *gin.Context) {
	var (
		resp struct {
			List []*service.UserInfo `json:"list"`
		}
		err error
	)
	defer func() {
		response.LogResponse(global.Logger, err)
		response.HandleResponse(ctx, err, resp)
	}()
	userId, _ := ctx.Get(global.UserIdKey)
	resp.List, err = services.FriendService.ListFriends(userId.(uint))
}

func DeleteFriend(ctx *gin.Context) {
	var (
		req  service.DeleteFriendReq
		resp service.DeleteFriendResp
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
	userId, _ := ctx.Get(global.UserIdKey)
	req.UserId = userId.(uint)
	resp, err = services.FriendService.DeleteFriend(req)
}

func registerFriendRouter(r *gin.Engine) {
	f := r.Group("/friends", middleware.Auth())
	{
		f.POST("/apply", ApplyFriend)
		f.PUT("/applyStatus", HandleApplyStatus)
		f.PUT("/applyQuestion", ApplyQuestion)
		f.PUT("/applyAnswer", ApplyAnswer)
		f.GET("/apply/list", ListApplication)
		f.GET("/list", ListFriend)
		f.DELETE("", DeleteFriend)
	}
}
