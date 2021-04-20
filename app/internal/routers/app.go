package routers

import (
	"os"
	"path/filepath"
	"together/app/internal/service"
	"together/app/pkg/ierr"
	"together/app/pkg/resp"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (a *App) GetDefaultInfo(c *gin.Context) {
	result := make(map[string]interface{}, 2)
	result["projectName"] = filepath.Base(os.Args[0])
	// 检测 grpc 链接
	s := service.New(c.Request.Context())
	r, err := s.SayHello("together")
	if err != nil {
		resp.Error(c, ierr.Grpc.WithDetails(err.Error()))
		return
	}
	result["GRPC"] = "OK"
	result["msg"] = r.Message
	resp.OK(c, result)
}
