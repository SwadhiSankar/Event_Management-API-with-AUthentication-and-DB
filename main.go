package main

import (

	// "example.com/main/models"
	"example.com/main.go/db"

	"example.com/main.go/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()

  server :=gin.Default()
  routes.RegisterRoutes(server)
 
  server.Run(":8080") 
}
