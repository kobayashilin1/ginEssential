package main

import (
	"github.com/kobayashilin1/ginEssential/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/oath/register", controller.Register)
	r.POST("/api/oath/login", controller.Login)

	return r
}