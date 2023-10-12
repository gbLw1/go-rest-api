package controllers

import (
	"go-rest-api/initializers"
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func PostsGetAll(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsGetById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsCreate(c *gin.Context) {
	var postBodyArgs struct {
		Title string
		Body  string
	}

	c.Bind(&postBodyArgs)

	post := models.Post{
		Title: postBodyArgs.Title,
		Body:  postBodyArgs.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var putBodyArgs struct {
		Title string
		Body  string
	}

	c.Bind(&putBodyArgs)

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	initializers.DB.Model(&post).Updates(&models.Post{
		Title: putBodyArgs.Title,
		Body:  putBodyArgs.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(204)
}
