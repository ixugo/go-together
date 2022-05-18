package routers

import (
	"os"
	"path/filepath"
	"together/app/internal/service"
	"together/app/pkg/ierr"
	"together/app/pkg/resp"

	"github.com/gin-gonic/gin"
)

type App struct {
	ser *service.Service
}

func (a *App) GetDefaultInfo(c *gin.Context) {
	result := make(map[string]interface{}, 2)
	result["projectName"] = filepath.Base(os.Args[0])
	// 检测 grpc 链接

	r, err := a.ser.SayHello(c.Request.Context(), "together")
	if err != nil {
		resp.Error(c, ierr.ErrUnknown.WithDetails(err.Error()))
		return
	}
	result["GRPC"] = "OK"
	result["msg"] = r.Message
	resp.OK(c, result)
}
