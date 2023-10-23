package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"

	"go-rest-api/pkg/initializers"
	"go-rest-api/pkg/models"
)

func validateRequestArgs(args models.Post) (*models.Post, error) {
	if args.Title == "" || args.Body == "" {
		return nil, errors.New("title and body cannot be empty")
	}

	return &args, nil
}

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
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
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

	// Parse request body
	if err := c.Bind(&postBodyArgs); err != nil {
		c.JSON(400, "Request body isn't a valid JSON")
		return
	}

	// Validate request body
	post, err := validateRequestArgs(models.Post{
		Title: postBodyArgs.Title,
		Body:  postBodyArgs.Body,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create post
	if err := initializers.DB.Create(post).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	// Check if post exists
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Parse request body
	var putBodyArgs struct {
		Title string
		Body  string
	}
	if err := c.Bind(&putBodyArgs); err != nil {
		c.JSON(400, "Request body isn't a valid JSON")
		return
	}

	// Validate request body
	postUpdate, err := validateRequestArgs(models.Post{
		Title: putBodyArgs.Title,
		Body:  putBodyArgs.Body,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update post
	initializers.DB.Model(&post).Updates(postUpdate)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(204)
}
