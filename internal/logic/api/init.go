package api

import (
	"github.com/gin-gonic/gin"
	"kim/internal/logic/dao"
	"kim/internal/logic/service"
)

var services *service.Service

func InitApi() {
	d := dao.InitDao()
	services = service.InitService(d)
}

func InitRouter(r *gin.Engine) {
	InitApi()

	registerUserRouter(r)
	registerFriendRouter(r)
}
