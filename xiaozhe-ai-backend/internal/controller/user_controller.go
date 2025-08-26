package controller

import (
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	ctx.JSON(200, "用户信息")
}
