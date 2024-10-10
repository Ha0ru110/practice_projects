package main

import (
	"fmt"
	"frstapi.com/eventorganisersystem/db"
	"frstapi.com/eventorganisersystem/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run("192.168.1.11:8080")
	if err != nil {
		fmt.Println("starting server failed")
		return
	}
}
