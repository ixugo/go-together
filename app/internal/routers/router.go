package routers

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger()) // 后面替换日志
	// service.New(context.TODO())

	var app App
	var blog Blog

	r.GET("/app", app.GetDefaultInfo)
	b := r.Group("/blog")
	{
		b.GET("/list", blog.GetList)
	}

	return r
}
