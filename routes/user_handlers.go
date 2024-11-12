package routes

import (
	"fmt"
	"net/http"
	"time"

	"example.com/go_app/models"
	"example.com/go_app/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	var id int64
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messae": "Could not parse the user",
		})
		return
	}
	user.Created_at = time.Now()
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Hash Password",
		})
		return
	}
	id, err = user.AddUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Create User",
		})
		return
	}
	user.ID = id
	context.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"event":   user,
	})
}

func getUsers(context *gin.Context) {
	events, err := models.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Server Error while parsing",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
		})
		return
	}
	fmt.Println(user)
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error while genrating token",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   token,
	})
}
