package controllers

import (
	"blog_sample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex handla la homepage
func HomeIndex(c *gin.Context) {

	models.GetPosts()
	//
	// c.HTML(http.StatusOK, "home-index.html", map[string]interface{}{
	// 	"Posts": posts,
	// })
	c.HTML(http.StatusOK, "home-index.html", nil)
}
