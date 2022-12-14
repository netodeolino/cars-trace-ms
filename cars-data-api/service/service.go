package service

import (
	"net/http"

	log "github.com/sirupsen/logrus"

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
	repo, err := repository.NewRepository(config)

	if err != nil {
		log.Error(err)
	}

	page := request.GeneratePageFromRequest(c)
	result, err := repo.GetAllData(&page)

	if err != nil {
		log.Error(err)
	}

	c.IndentedJSON(http.StatusOK, result)
}
