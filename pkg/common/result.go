package common

import (
	"fmt"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(data any) Result {
	return Result{
		Code: 200,
		Msg:  "OK",
		Data: data,
	}
}

func Bad(msg string) Result {
	return Result{
		Code: 400,
		Msg:  msg,
		Data: nil,
	}
}

func InternalError(msg string) Result {
	return Result{
		Code: 500,
		Msg:  msg,
		Data: nil,
	}
}

func NoAuth(msg string) Result {
	return Result{
		Code: 403,
		Msg:  fmt.Sprintf("Unauthorized: %s", msg),
		Data: nil,
	}
}

func NotFound() Result {
	return Result{
		Code: 404,
		Msg:  "Not Found",
		Data: nil,
	}
}
