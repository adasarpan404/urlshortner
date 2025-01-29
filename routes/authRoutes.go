package routes

import (
	"github.com/adasarpan404/urlshortner/controller"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", controller.Signup)
		authGroup.POST("/login", controller.Login)
	}
}
