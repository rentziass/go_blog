package router

import (
	"blog_sample/controllers"

	"github.com/gin-gonic/gin"
)

func registerPostsRoutes(r *gin.Engine) {
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts/:id", controllers.PostsShow)
}
