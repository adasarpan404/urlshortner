package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	SetupAuthRoutes(r)
	SetupURLRoutes(r)
	SetupUserRoutes(r)
	return r
}
