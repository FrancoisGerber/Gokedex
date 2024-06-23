package middleware

import (
	"Gokedex/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	valid, userId, err := utils.ValidateJWT(token)

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !valid {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// loggedInUserId := context.GetInt64("userId")
	context.Set("userId", userId)
	context.Next()
}
