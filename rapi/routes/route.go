package routes

import (
	"example.com/rapi/middlewares"
	"github.com/gin-gonic/gin"
)
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) 
	server.GET("/events/:id", getEvent) 
	authenticated:=server.Group(("/"))
	authenticated.Use()
	authenticated.POST("/events",middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
     authenticated.DELETE("/events/:id", deleteEvent)
	 authenticated.POST("events/:id/register",registerForEvent)
	 authenticated.DELETE("events/:id/register",DeleteAnEvent)
	server.POST("/signup",signup)
	server.POST("/login",login)
	
}