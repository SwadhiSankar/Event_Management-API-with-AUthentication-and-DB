package main

import (
	"net/http"

	// "example.com/main/models"
	"example.com/main.go/db"
	"example.com/main.go/models"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()

  server :=gin.Default()

  server.GET("/EVENTS",getEvents)  //GET, POST, PUT, PATCH, DELETE
  server.POST("/events",createEvent)
  server.Run(":8080") //starts listening incoming request
}

func getEvents(context *gin.Context ) {
 events := models.GetAllEvent()
 context.JSON(http.StatusOK,events)
}

func createEvent(context *gin.Context){
 var event models.Event

 err  := context.ShouldBindJSON(&event)

 if err !=nil {
 context.JSON(http.StatusBadRequest,gin.H{"message": "Missing required fields data"})
 return
 }
 event.ID = 1
 event.UserID = 1
 context.JSON(http.StatusCreated,gin.H{"message":"Eventcreated","event": event})
}