package main

import (
	"fmt"
	"github.com/muhammad-usman13/jwt-golang/handlers"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.POST("/login", handlers.LoginHandler)
	router.POST("/refresh", handlers.RefreshHandler)
	router.GET("/protected", handlers.ProtectedHandler)
	
	fmt.Println("Starting the server")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
}
