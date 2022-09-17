package routes

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/netodeolino/cars-trace-ms/cars-data-api/service"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Requesting API from IP: %s\n", c.ClientIP())
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

	router.Run()
}
