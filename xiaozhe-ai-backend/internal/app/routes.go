package app

import (
	"xiaozhe-ai-backend/internal/common"
	"xiaozhe-ai-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/health", common.HealthCheck)

	user := r.Group("/api/user")
	{
		user.GET("/getUserInfo", controller.GetUserInfo)
	}

	return r
}
