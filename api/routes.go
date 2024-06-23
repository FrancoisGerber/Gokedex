package api

import (
	"Gokedex/api/auth"
	"Gokedex/api/controllers"
	"Gokedex/api/middleware"

	"github.com/gin-gonic/gin"
)

// Setup API route groups and endpoints
func Routes(router *gin.Engine) {
	//Open endpoints
	authGroup := router.Group("/auth")

	//Closed endpoint (needs authentication)
	apiGroup := router.Group("/api", middleware.Authentication)

	//Add all the endpoints
	auth.SetupEndpoints(authGroup)

	users.SetupEndpoints(apiGroup)
}
