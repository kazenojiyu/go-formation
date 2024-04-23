package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"projet.com/rest-api/utils"
)

func Autenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized (token empty)"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized (bad token)"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
