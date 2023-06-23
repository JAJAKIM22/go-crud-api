package main

import (
	"github.com/jajakim22/go-crud/initializers"
	"github.com/jajakim22/go-crud/models"
)

func init() {
initializers.LoadEnvVariables()
initializers.ConnectToDB()

}


func  main() {
	initializers.DB.AutoMigrate(&models.Post{})
}