package main

import (
	"go-rest-api/pkg/controllers"
	"go-rest-api/pkg/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.Introduction)
	r.GET("/posts", controllers.PostsGetAll)
	r.GET("/posts/:id", controllers.PostsGetById)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
