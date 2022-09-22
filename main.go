package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salimkun/sample-crud-go/service"
)

func main() {
	router := gin.Default()

	version1 := router.Group("/api/v1")

	version1.POST("/user", service.RegisterUser)
	version1.GET("/user", service.ListUser)
	version1.PATCH("/user", service.UpdateUser)
	version1.GET("/user/:userID", service.GetUserByID)
	version1.DELETE("/user/:userID", service.DeleteUser)

	router.Run("localhost:8080")
}
