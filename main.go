package main

import (
	"net/http"
	"strconv"

	// "example.com/main/models"
	"example.com/main.go/db"
	"example.com/main.go/models"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()

  server :=gin.Default()

  server.GET("/EVENTS",getEvents)  //GET, POST, PUT, PATCH, DELETE
  server.GET("/events/:id",getEvent)// id is dynamic path 
  server.POST("/events",createEvent) 
  server.Run(":8080") //starts listening incoming request
}

func getEvents(context *gin.Context ) {
 events,err := models.GetAllEvent()
 if err!=nil{
	context.JSON(http.StatusInternalServerError,gin.H{"Message":"Issue while fetching Try again"})
	return
 }
 context.JSON(http.StatusOK,events)
}


func getEvent(context *gin.Context){
   eventId, err := strconv.ParseInt(context.Param("id"),10,64)  //10 decimal system //64 bit size
   if err != nil{
	context.JSON(http.StatusBadRequest,gin.H{"message": "Cannot able to parse event ID"})
	return
   }
   event, err := models.GetEventById(eventId)
   if err != nil{
	context.JSON(http.StatusInternalServerError,gin.H{"message": "not able to fetch the event"})
	return
   }
   context.JSON(http.StatusOK,event)

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
 err  = event.Save()
 if err!=nil{
	context.JSON(http.StatusInternalServerError,err.Error())
	return
 }
 context.JSON(http.StatusCreated,gin.H{"message":"Eventcreated","event": event})
}