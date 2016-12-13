package controllers

import (
	"net/http"
	"strconv"

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

func PostsShow(c *gin.Context) {
	if c.Param("id") == "new" {
		PostsNew(c)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 0, 32)
	if err != nil {
		c.String(http.StatusBadRequest, "ID Format not valid")
		return
	}

	post, err := models.GetPost(id)
	if err != nil {
		c.String(http.StatusNotFound, "No post found with given id")
	}

	// c.HTML(http.StatusOK, "posts-show.html", map[string]interface{}{
	// 	"Title": post.Title,
	// 	"Body":  post.Body,
	// })
	c.HTML(http.StatusOK, "posts-show.html", post)
}
