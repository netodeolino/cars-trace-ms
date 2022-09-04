package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func Start() {
	router := gin.Default()
	router.GET("/", index)
	router.Run("localhost:8080")
}
