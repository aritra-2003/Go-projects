package routes

import (
	"example.com/financetracker/events"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", events.Register)
	r.POST("/login", events.Login)
}