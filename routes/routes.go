package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", welcome)
	server.POST("/emi", calculateLoanEMI)
}
