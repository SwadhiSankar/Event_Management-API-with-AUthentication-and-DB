package routes

import (
	"example.com/main.go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/EVENTS", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // id is dynamic path
	    
	authenticated := server.Group("/")
	authenticated.Use( middlewares.Authentication)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
    authenticated.DELETE("/events/:id/register")
	
	server.POST("/signup",signup)
	server.POST("/login",login)
}