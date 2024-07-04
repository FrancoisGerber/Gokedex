package pokemon

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
	pokemonGroup := utils.SetupRoute("pokemon", api, utils.Handler{
		"GetAll": GetAll,
		"Get":    Get,
		"Post":   Post,
		"Put":    Put,
		"Delete": Delete,
	})
	_ = pokemonGroup
}

// Get All Pokemon godoc
// @Summary Get all the Pokemon in the pokedex
// @Schemes
// @Description Get all the Pokemon in the pokedex
// @Tags Pokemon
// @Accept json
// @Produce json
// @Success 200 {array} models.Pokemon
// @Router /pokemon [get]
// @Security ApiKeyAuth
func GetAll(context *gin.Context) {
	var pokemon = interfaces.IModel[models.Pokemon](&models.Pokemon{})
	data, err := pokemon.GetAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	if data == nil {
		data = []models.Pokemon{}
	}

	context.JSON(http.StatusOK, data)
}

func Get(context *gin.Context) {
	var pokemon = interfaces.IModel[models.Pokemon](&models.Pokemon{})
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	err = pokemon.Get(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, pokemon)
}

func Post(context *gin.Context) {
	var pokemon = interfaces.IModel[models.Pokemon](&models.Pokemon{})
	err := context.ShouldBindJSON(&pokemon)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	err = pokemon.Create()

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, pokemon)
}

func Put(context *gin.Context) {
	var pokemon = interfaces.IModel[models.Pokemon](&models.Pokemon{})
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	err = context.ShouldBindBodyWithJSON(&pokemon)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	rowsChanged, err := pokemon.Update(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, gin.H{"model": pokemon, "rows": rowsChanged})
}

func Delete(context *gin.Context) {
	var pokemon = interfaces.IModel[models.Pokemon](&models.Pokemon{})
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.JsonError(err, http.StatusBadRequest))
		return
	}

	deleted, err := pokemon.Delete(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.JsonError(err, http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, gin.H{"Deleted": deleted})
}
