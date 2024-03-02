package common

import (
	"fmt"
)

const (
	OK                    = 200
	BAD_REQUEST           = 400
	UNAUTHORIZED          = 401
	FORBIDDEN             = 403
	NOT_FOUND             = 404
	INTERNAL_SERVER_ERROR = 500
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(data any) Result {
	return Result{
		Code: OK,
		Msg:  "OK",
		Data: data,
	}
}

func Bad(msg string) Result {
	return Result{
		Code: BAD_REQUEST,
		Msg:  msg,
		Data: nil,
	}
}

func InternalError(msg string) Result {
	return Result{
		Code: INTERNAL_SERVER_ERROR,
		Msg:  msg,
		Data: nil,
	}
}

func NoAuth(msg string) Result {
	return Result{
		Code: UNAUTHORIZED,
		Msg:  fmt.Sprintf("Unauthorized: %s", msg),
		Data: nil,
	}
}

func Forbidden(msg string) Result {
	return Result{
		Code: FORBIDDEN,
		Msg:  fmt.Sprintf("forbidden: %s", msg),
		Data: nil,
	}
}

func NotFound() Result {
	return Result{
		Code: NOT_FOUND,
		Msg:  "Not Found",
		Data: nil,
	}
}
