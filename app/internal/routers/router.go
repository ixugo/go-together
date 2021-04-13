package routers

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) // 后面替换日志

	var app App
	var blog Blog

	v1 := r.Group("/v1")
	{
		v1.GET("/app", app.GetDefaultInfo)
		b := v1.Group("/blog")
		{
			b.GET("/list", blog.GetList)
		}
	}

	return r
}
