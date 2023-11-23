package main

// import "fmt"

import (
	"github.com/gin-gonic/gin"
	"github.com/inabatron/go_crud/controllers"
	"github.com/inabatron/go_crud/inits"
)

func init() {
	inits.LoadEnvVars()
	inits.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)

	r.GET("/posts", controllers.PostsIndex)

	r.GET("/posts/:id", controllers.PostsShow)

	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
