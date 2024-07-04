package users

import (
	"Gokedex/data/interfaces"
	"Gokedex/data/models"
	"Gokedex/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Setup the endpoints for this controller's handlers
func SetupEndpoints(api *gin.RouterGroup) {
	userGroup := utils.SetupRoute("users", api, utils.Handler{
		"GetAll": GetAll,
		"Get":    Get,
		"Post":   Post,
		"Put":    Put,
		"Delete": Delete,
	})
	_ = userGroup
}

func GetAll(context *gin.Context) {
	var user = interfaces.IModel[models.User](&models.User{})
	data, err := user.GetAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, data)
}

func Get(context *gin.Context) {
	var user = interfaces.IModel[models.User](&models.User{})
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	err = user.Get(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, user)
}

func Post(context *gin.Context) {
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

func Put(context *gin.Context) {
	var user = interfaces.IModel[models.User](&models.User{})
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	err = context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	rowsChanged, err := user.Update(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, gin.H{"model": user, "rows": rowsChanged})
}

func Delete(context *gin.Context) {
	var user = interfaces.IModel[models.User](&models.User{})
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	deleted, err := user.Delete(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, gin.H{"Deleted": deleted})
}
