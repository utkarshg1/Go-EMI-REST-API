package main

import (
	"example.com/emi-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
