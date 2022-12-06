package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`    // 错误编码
	Msg     string   `json:"msg"`     // 错误消息
	Details []string `json:"details"` // 详细消息
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d,已存在,请更换一个", code))
	}

	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码code: %d, 错误信息msg: %s, 详细消息details: %s", e.Code, e.Msg, e.Details)
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) Detail() []string {
	return e.Details
}

func (e *Error) WithDetails(details ...string) *Error {
	e.Details = []string{}
	for _, d := range details {
		e.Details = append(e.Details, d)
	}
	return e
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
