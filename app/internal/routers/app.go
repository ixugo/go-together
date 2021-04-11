package routers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (a *App) GetDefaultInfo(c *gin.Context) {
	// 返回执行文件名
	c.String(http.StatusOK, filepath.Base(os.Args[0]))
}
