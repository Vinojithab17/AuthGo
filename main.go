package main

import (
	"example.com/go_app/db"
	"example.com/go_app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //http://localhot:8080
}
