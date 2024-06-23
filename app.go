package main

import (
	"Gokedex/api"
	"Gokedex/data"

	"github.com/gin-gonic/gin"
)

// Pokedex
func main() {
	router := gin.New()

	//Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Routes
	api.Routes(router)

	//Setup Database
	data.InitDB()
	defer data.CloseDB()

	router.Run(":8051")
}
