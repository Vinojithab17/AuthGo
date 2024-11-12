package routes

import (
	"example.com/go_app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/event", getEvents)
	server.GET("/users", getUsers)
	server.GET("/event/:id", getEventByID)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/event", createEvent)
	authenticate.PUT("/event/:id", updateEvent)
	authenticate.DELETE("/event/:id", deleteEvent)

	server.POST("/signup", createUser)
	server.POST("/login", login)
}
