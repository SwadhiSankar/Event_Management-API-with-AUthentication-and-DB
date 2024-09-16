package routes

import (
	"net/http"
	"strconv"

	"example.com/main.go/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
   userId := context.GetInt64("userId")
   eventId, err :=  strconv.ParseInt(context.Param("id"), 10, 64) //10 decimal system //64 bit size
   if err != nil {
	   context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot able to parse event ID"})
	   return
   }

   event, err := models.GetEventById(eventId)
   if err !=nil{
	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
	return
   }

  err = event.Register(userId)

  if err !=nil{
	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register an event"})
	return
   }
  
   context.JSON(http.StatusCreated, gin.H{"message": "Registered Successfully"})
	

}

func cancelRegistration(context *gin.Context) {

}