package middlewares

import (
	"fmt"
	"net/http"

	"example.com/go_app/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to perform this action",
		})
		return
	}
	user_id, err := utils.VerifyToken(token)
	fmt.Println("this is the user id", user_id)
	if err != nil {
		// panic(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}
	context.Set("user_id", user_id)
	context.Next()
}
