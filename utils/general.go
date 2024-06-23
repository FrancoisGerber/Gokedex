package utils

import (
	"github.com/gin-gonic/gin"
)

// Shorthand for a Gin endpoint handler
type Handler map[string]func(*gin.Context)

// The type all models must be converted into to interact with the db
type DataModel map[string]any

// Binds CRUD endpoints to their handler functions
func SetupRoute(name string, group *gin.RouterGroup, handlers Handler) *gin.RouterGroup {
	subGroup := group.Group("/" + name)
	{
		subGroup.GET("/", handlers["GetAll"])
		subGroup.GET("/:id", handlers["Get"])
		subGroup.POST("/", handlers["Post"])
		subGroup.PUT("/:id", handlers["Put"])
		subGroup.DELETE("/:id", handlers["Delete"])
	}

	return subGroup
}

// Create a JSON error object
func JsonError(err error, status int) gin.H {
	return gin.H{
		"Error": err.Error(),
		"Code":  status,
	}
}
