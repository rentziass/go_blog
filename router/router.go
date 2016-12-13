package router

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")

	registerHomeRoutes(r)
	registerPostsRoutes(r)

	r.Static("/public", "./public")

	return r
}
