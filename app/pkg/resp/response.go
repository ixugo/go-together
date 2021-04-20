package resp

import (
	"net/http"
	"together/app/pkg/ierr"

	"github.com/gin-gonic/gin"
)

// OK 通用成功返回
func OK(c *gin.Context, bean interface{}) {
	c.JSON(http.StatusOK, bean)
}

// Error 通用错误返回
func Error(c *gin.Context, err *ierr.Error) {
	r := gin.H{"code": err.Code(), "msg": err.Msg()}
	d := err.Details()
	if len(d) > 0 {
		r["details"] = d
	}
	c.JSON(err.StatusCode(), r)
}
