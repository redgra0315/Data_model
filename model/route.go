package main

import (
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	gin.ForceConsoleColor()

	// 路由
	v1 := r.Group("/api")
	{
		v1.POST("create", CreateData)
		v1.GET("get/:id", GetData)
	}
	return r
}
