package routers

import (
	"net/http"
	"os"
	"path/filepath"
	"together/app/internal/service"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (a *App) GetDefaultInfo(c *gin.Context) {
	// 返回执行文件名
	c.String(http.StatusOK, filepath.Base(os.Args[0]))
	// 检测 grpc 链接
	s := service.New(c.Request.Context())
	r, err := s.SayHello("together")
	c.String(http.StatusOK, "\nGRPC \n")
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, r.Message)
}
