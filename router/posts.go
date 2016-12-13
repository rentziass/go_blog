package router

import (
	"blog_sample/controllers"

	"github.com/gin-gonic/gin"
)

func registerPostsRoutes(r *gin.Engine) {
	r.GET("/posts/new", controllers.PostsNew)
}
