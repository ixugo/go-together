package resp

import (
	"net/http"
)

// OK 通用成功返回
func OK(c ResponseWriter, bean interface{}) {
	c.JSON(http.StatusOK, bean)
}

// Errorer ...
type Errorer interface {
	Code() int
	HTTPCode() int
	Message() string
	Details() []string
}

// ResponseWriter ...
type ResponseWriter interface {
	JSON(code int, obj interface{})
	File(filepath string)
}

// Error 通用错误返回
func Error(c ResponseWriter, err Errorer) {
	r := map[string]interface{}{"code": err.Code(), "msg": err.Message()}
	d := err.Details()
	if len(d) > 0 {
		r["details"] = d
	}
	c.JSON(err.HTTPCode(), r)
}
