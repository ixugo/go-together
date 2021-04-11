package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (a *App) GetDefaultInfo(c *gin.Context) {

	// 请返回执行文件名
	c.String(http.StatusOK, "........")

}
