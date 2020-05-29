package main

import (
	"ginEssential/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/oath/register", controller.Register)

	return r
}
