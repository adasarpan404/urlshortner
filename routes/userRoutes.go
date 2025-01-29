package routes

import (
	"github.com/adasarpan404/urlshortner/controller"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/profile", controller.GetUserProfile)
		userGroup.PUT("/update", controller.UpdateUserProfile)
	}
}
