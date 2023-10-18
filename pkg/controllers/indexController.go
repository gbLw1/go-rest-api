package controllers

import "github.com/gin-gonic/gin"

type endpointInfo struct {
	Path        string
	Description string
}

func Introduction(c *gin.Context) {
	endpoints := []endpointInfo{
		{
			Path:        "[GET] ~/",
			Description: "Introduction to all endpoints available in the API",
		},
		{
			Path:        "[GET] ~/posts",
			Description: "Get all posts",
		},
		{
			Path:        "[GET] ~/posts/{id}",
			Description: "Get a post by id",
		},
		{
			Path:        "[POST] ~/posts",
			Description: "Create a new post",
		},
		{
			Path:        "[PUT] ~/posts/{id}",
			Description: "Update a existing post by id",
		},
		{
			Path:        "[DELETE] ~/posts{id}",
			Description: "Delete a post by id",
		},
	}

	c.JSON(200, gin.H{
		"API":       "Go Rest API is running...",
		"endpoints": endpoints,
	})
}
