package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type App struct{}

func (a *App) GetDefaultInfo(c *gin.Context) {

	// 返回执行文件名
	c.String(http.StatusOK, filepath.Base(os.Args[0]))

}
