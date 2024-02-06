package controller

import "github.com/gin-gonic/gin"

func Ok(data any) gin.H {
	return gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	}
}

func Bad(msg string) gin.H {
	return gin.H{
		"code": 400,
		"msg":  msg,
		"data": nil,
	}
}

func Error(msg string) gin.H {
	return gin.H{
		"code": 500,
		"msg":  msg,
		"data": nil,
	}
}

func NotFound(msg string) gin.H {
	return gin.H{
		"code": 404,
		"msg":  msg,
		"data": nil,
	}
}
