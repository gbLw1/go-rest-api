package main

import (
	"github.com/gin-gonic/gin"

	"go-rest-api/pkg/controllers"
	"go-rest-api/pkg/initializers"
	"go-rest-api/pkg/middleware"
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
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.PostsUpdate)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.PostsDelete)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.Run()
}
