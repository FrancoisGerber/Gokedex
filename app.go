package main

import (
	"Gokedex/api"
	"Gokedex/data"

	docs "Gokedex/api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title           Gokedex
// @version         1.0
// @description     A Go Web API to store and manage pokemon, aka "Pokedex".

// @contact.name   Francois Gerber

// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	//Setup Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8051")
}
