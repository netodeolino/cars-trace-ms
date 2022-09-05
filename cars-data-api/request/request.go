package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/netodeolino/cars-trace-ms/cars-data-core/models"
)

func GeneratePageFromRequest(c *gin.Context) models.Page[models.CarModel] {
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()

loop:
	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break loop
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break loop
		case "sort":
			sort = queryValue
			break loop
		}
	}

	return models.Page[models.CarModel]{Limit: limit, Page: page, Sort: sort, Data: models.CarModel{}}
}
