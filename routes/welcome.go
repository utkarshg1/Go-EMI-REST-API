package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome to EMI Calculator API!",
	})
}
