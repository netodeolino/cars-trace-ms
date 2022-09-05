package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func Teste() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Request with param %s", c.Query("page"))
	}
}

func Start() {
	router := gin.Default()
	router.GET("/", index)

	api := router.Group("/api")

	api.Use(Teste())
	{
		api.GET("/", index)
	}

	router.Run("localhost:8080")
}
