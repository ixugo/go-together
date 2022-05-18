package routers

import (
	"together/app/internal/http/middleware"
	"together/app/internal/service"
	"together/configs"

	"github.com/gin-gonic/gin"
)

func New(cfg *configs.AppServer) *gin.Engine {
	r := gin.New()

	ser := service.New(cfg)

	r.Use(gin.Logger()) // 后面替换日志
	r.Use(middleware.Header())

	newApp(r, ser)
	newBlog(r, ser)

	return r
}

func newApp(r *gin.Engine, ser *service.Service) {
	app := App{ser: ser}
	r.GET("/app", app.GetDefaultInfo)
}

func newBlog(r *gin.Engine, ser *service.Service) {
	blog := Blog{ser: ser}
	r.GET("/blog", blog.GetList)
}
