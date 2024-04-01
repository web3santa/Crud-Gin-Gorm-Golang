package main

import (
	"github.com/web3santa/Crud-Gin-Gorm-Golang/initializers"
	"github.com/web3santa/Crud-Gin-Gorm-Golang/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
