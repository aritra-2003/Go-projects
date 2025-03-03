package main

import (
	"example.com/rapi/db"
	"example.com/rapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
    db.InitDB()
	// Create an event
    routes.RegisterRoutes(server)
	server.Run(":8080") // Listen on port 8080
}

