package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/EVENTS", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // id is dynamic path
	server.POST("/events", createEvent)
	server.PUT("/events/:id",updateEvent) //update req
}