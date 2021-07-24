// 自定义错误

package ierr

import (
	"fmt"
	"net/http"
)

var (
	Success                = NewError(0, "成功")
	BadRequest             = NewError(40000, "请求参数错误")
	NotFound               = NewError(40001, "资源找不到")
	UnauthorizedTokenError = NewError(40002, "鉴权失败，Token错误")
	Server                 = NewError(50000, "服务内部错误")
	Grpc                   = NewError(50001, "GRPC 连接失败")
)

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

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(args ...string) *Error {
	newErr := *e
	newErr.details = make([]string, 0, len(args))
	newErr.details = append(newErr.details, args...)
	return &newErr
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case Server.Code():
		return http.StatusInternalServerError
	case BadRequest.Code():
		return http.StatusBadRequest
	case UnauthorizedTokenError.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
