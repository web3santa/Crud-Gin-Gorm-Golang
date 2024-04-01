package main

import (
	"github.com/gin-gonic/gin"
	"github.com/web3santa/Crud-Gin-Gorm-Golang/controllers"
	"github.com/web3santa/Crud-Gin-Gorm-Golang/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostsCreate)
	r.GET("/get", controllers.PostsIndex)
	r.GET("/get/:id", controllers.PostsShow)
	r.DELETE("/get/:id", controllers.PostsDelete)
	r.PUT("/get/:id", controllers.PostsUpdate)

	r.Run(":3000")
}
