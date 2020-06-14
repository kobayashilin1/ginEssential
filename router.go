package main

import (
	"github.com/kobayashilin1/ginEssential/controller"
	"github.com/kobayashilin1/ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(),controller.Info)
	//使用AuthMiddleware实现用户认证，保护数据

	return r
}