package routers

import (
	"net/http"
	"together/app/internal/service"

	"github.com/gin-gonic/gin"
)

type Blog struct{}

func (b *Blog) GetList(c *gin.Context) {
	url := c.Query("url")
	// 处理链接
	s := service.New(c.Request.Context())
	data, err := s.GetList(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
