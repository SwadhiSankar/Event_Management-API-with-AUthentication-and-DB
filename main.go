package main 


import "github.com/gin-gonic/gin"
import "net/http"

func main(){
  server :=gin.Default()

  server.GET("/EVENTS",getEvents)  //GET, POST, PUT, PATCH, DELETE

  server.Run(":8080") //starts listening incoming request
}

func getEvents(context *gin.Context ) {
 context.JSON(http.StatusOK,gin.H{"message": "Hello"})
}