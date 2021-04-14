package routers

import (
	"together/app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger()) // 后面替换日志
	r.Use(middleware.Header())

	var app App
	var blog Blog

	r.GET("/app", app.GetDefaultInfo)
	r.GET("/blog", blog.GetList)

	return r
}
