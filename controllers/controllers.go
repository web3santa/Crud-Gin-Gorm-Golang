package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/web3santa/Crud-Gin-Gorm-Golang/initializers"
	"github.com/web3santa/Crud-Gin-Gorm-Golang/models"
)

func PostsCreate(ctx *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	// return it

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	// get the post
	// Get all records
	var posts []models.Post
	initializers.DB.Find(&posts)

	// SELECT * FROM users;

	ctx.JSON(200, gin.H{
		"posts": posts,
	}) // returns error
	// respond with them
}

func PostsShow(ctx *gin.Context) {

	// get id off url
	id := ctx.Param("id")
	// get the post
	// Get all records
	var post models.Post
	initializers.DB.First(&post, id)

	// SELECT * FROM users;

	ctx.JSON(200, gin.H{
		"posts": post,
	}) // returns error
	// respond with them
}

func PostsDelete(ctx *gin.Context) {

	// get id off url
	id := ctx.Param("id")
	// get the post
	// Get all records
	initializers.DB.Delete(&models.Post{}, id)

	// SELECT * FROM users;

	ctx.Status(200)
	// respond with them
}

func PostsUpdate(ctx *gin.Context) {

	// get id off url
	id := ctx.Param("id")
	// get the post
	// Get all records
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)
	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// SELECT * FROM users;

	ctx.JSON(200, gin.H{
		"posts": post,
	}) // returns error
	// respond with them
}
