package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsNew(c *gin.Context) {
	c.HTML(http.StatusOK, "posts-new.html", nil)
}
