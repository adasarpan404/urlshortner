package routes

import (
	"github.com/adasarpan404/urlshortner/controller"
	"github.com/gin-gonic/gin"
)

func SetupURLRoutes(r *gin.Engine) {
	r.POST("/shorten", controller.ShortenUrl)
	r.GET("/:shortCode", controller.RedirectUrl)
}
