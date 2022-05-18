// 自定义错误

package ierr

import (
	"fmt"
	"net/http"
)

// 业务常用错误
var (
	ErrUnknown           = NewError(10101, "未知错误")
	BadRequest           = NewError(10102, "请求参数有误")
	ErrDB                = NewError(10201, "数据库发生错误")
	ErrUnauthorizedToken = NewError(10301, "TOKEN 验证失败")
	ErrJSON              = NewError(10401, "JSON 编解码出错")
)

// Error ...
type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

// NewError 创建自定义错误
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// Code ..
func (e *Error) Code() int {
	return e.code
}

// Message ..
func (e *Error) Message() string {
	return e.msg
}

// Details 错误
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 错误详情
func (e *Error) WithDetails(args ...string) *Error {
	newErr := *e
	newErr.details = make([]string, 0, len(args))
	newErr.details = append(newErr.details, args...)
	return &newErr
}

// HTTPCode http status code
func (e *Error) HTTPCode() int {
	switch e.Code() {
	case 0:
		return http.StatusOK
	case ErrUnauthorizedToken.code:
		return http.StatusUnauthorized
	}
	return http.StatusBadRequest
}
