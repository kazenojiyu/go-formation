package main

import (
	"github.com/gin-gonic/gin"
	"projet.com/rest-api/db"
	"projet.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
