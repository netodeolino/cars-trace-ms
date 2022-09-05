package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/netodeolino/cars-trace-ms/cars-data-api/service"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Requesting API from IP: %s\n", c.ClientIP())
	}
}

func Start() {
	router := gin.Default()
	service := &service.Service{}

	router.GET("/", service.Index)

	api := router.Group("/api")
	api.Use(middleware())
	{
		api.GET("/", service.GetAllData)
	}

	router.Run("localhost:8080")
}
