package main

import (
	"example.com/financetracker/database"
	"example.com/financetracker/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	database.InitDB()
	r := gin.Default()
	routes.AuthRoutes(r)
	routes.TransactionRoutes(r)
	r.Run(":8080")
	}

