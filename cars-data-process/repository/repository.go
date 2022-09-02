package repository

import (
	"cars-data/process/config"
	"cars-data/process/models"

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

func (repo *Repository) SaveCsvData(wg *sync.WaitGroup, data []models.CarEntity) {
	defer wg.Done()
	repo.DB.CreateInBatches(data, 100)
}
