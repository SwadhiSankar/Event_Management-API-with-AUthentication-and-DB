package routes

import (
	"net/http"

	"example.com/main.go/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields data"})
		return
	}
    
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"User created Successfully"})
}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields data"})
		return
	}
    
	err = user.ValidateCredentials()

	if err != nil{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"UnAuthorized user"})
		return
	}

   context.JSON(http.StatusOK,gin.H{"message":"Successfully logged in."})
}