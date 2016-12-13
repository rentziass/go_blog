package router

import (
	"blog_sample/controllers"

	"github.com/gin-gonic/gin"
)

func registerHomeRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomeIndex)
}
