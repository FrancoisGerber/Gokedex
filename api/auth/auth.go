package auth

import (
	"Gokedex/data/interfaces"
	"Gokedex/data/models"
	"Gokedex/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup the endpoints for this controller's handlers
func SetupEndpoints(api *gin.RouterGroup) {
	api.POST("/register", Register)
	api.POST("/login", Login)
}

func Login(context *gin.Context) {
	var user = interfaces.IModel[models.User](&models.User{})
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	u := user.(*models.User)

	valid, token, err := u.ValidateCredentials()

	context.JSON(http.StatusOK, gin.H{
		"Valid": valid,
		"Token": token,
	})
}

func Register(context *gin.Context) {
	var user = interfaces.IModel[models.User](&models.User{})
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	err = user.Create()

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, user)
}
