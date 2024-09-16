package routes

import (
	"net/http"

	"example.com/main.go/models"
	"example.com/main.go/utils"
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

	token, err :=utils.GenerateToken(user.Email,user.ID)

	if err !=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not authenticate user"})
		return
	}
   context.JSON(http.StatusOK,gin.H{"message":"Successfully logged in.","token":token})
}