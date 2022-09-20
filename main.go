package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salimkun/sample-crud-go/service"
)

func main() {
	router := gin.Default()

	version1 := router.Group("/api/v1")

	version1.POST("/user", service.RegisterUser)
	router.Run("localhost:8080")
}
