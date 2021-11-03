package routers

import (
	"together/app/internal/service"
	"together/app/pkg/ierr"
	"together/app/pkg/resp"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ser *service.Service
}

func (b *Blog) GetList(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		resp.Error(c, ierr.BadRequest.WithDetails("url cannot be empty"))
		return
	}
	// 处理链接
	data, err := b.ser.GetList(c.Request.Context(), url)
	if err != nil {
		resp.Error(c, ierr.GetBlog.WithDetails(err.Error()))
		return
	}
	resp.OK(c, data)
}
