package repository

import (
	"github.com/netodeolino/cars-trace-ms/cars-data-core/config"
	"github.com/netodeolino/cars-trace-ms/cars-data-core/models"

	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo *Repository) InitializeDatabase(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Database,
		config.DB.Port,
		config.DB.SSLmode,
		config.DB.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	repo.DB = models.MigrateModels(db)
}

func (repo *Repository) SaveCsvData(wg *sync.WaitGroup, data []models.CarModel) {
	defer wg.Done()
	repo.DB.CreateInBatches(data, 100)
}

func (repo *Repository) getData(page *models.Page[models.CarModel]) (*models.Page[[]models.CarModel], error) {
	offset := (page.Page - 1) * page.Limit
	queryBuider := repo.DB.Limit(page.Limit).Offset(offset).Order(page.Sort)
	newPage := models.Page[[]models.CarModel]{Limit: page.Limit, Page: page.Page, Sort: page.Sort, Data: nil}

	result := queryBuider.Model(&models.CarModel{}).Where(page.Data).Find(&newPage.Data)

	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}

	return &newPage, nil
}
