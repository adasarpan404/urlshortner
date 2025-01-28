package routes

import (
	"github.com/adasarpan404/urlshortner/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/shorten", controller.ShortenUrl)
	r.GET("/:shortCode", controller.GET)

	return r
}
