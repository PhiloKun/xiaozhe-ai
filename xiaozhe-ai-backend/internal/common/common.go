package common

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功
func Success(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

// Fail 失败
func Fail(msg string) *Response {
	return &Response{
		Code: -1,
		Msg:  msg,
		Data: nil,
	}
}

// HealthCheck 健康检查
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}
