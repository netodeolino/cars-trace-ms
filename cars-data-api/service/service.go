package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netodeolino/cars-trace-ms/cars-data-api/request"
	"github.com/netodeolino/cars-trace-ms/cars-data-core/config"
	"github.com/netodeolino/cars-trace-ms/cars-data-core/repository"
)

type Service struct{}

func (service *Service) Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func (service *Service) GetAllData(c *gin.Context) {
	config := config.GetConfig()
	repo := repository.NewRepository(config)
	page := request.GeneratePageFromRequest(c)
	result := repo.GetAllData(page)

	c.IndentedJSON(http.StatusOK, result)
}
