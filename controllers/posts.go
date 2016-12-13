package controllers

import (
	"net/http"

	"blog_sample/models"

	"github.com/gin-gonic/gin"
)

func PostsNew(c *gin.Context) {
	c.HTML(http.StatusOK, "posts-new.html", nil)
}

func PostsCreate(c *gin.Context) {
	post := models.Post{}
	post.SetTitle(c.PostForm("title"))
	post.SetBody(c.PostForm("body"))

	post.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}
