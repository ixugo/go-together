package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) // 后面替换日志

	var app App

	v1 := r.Group("/api/v1")
	{
		v1.GET("/app", app.GetDefaultInfo)

	}
	return r
}
