package controllers

import (
	"blog_sample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex handla la homepage
func HomeIndex(c *gin.Context) {

	posts, _ := models.GetPosts()

	c.HTML(http.StatusOK, "home-index.html", map[string]interface{}{
		"Posts": posts,
	})
}
