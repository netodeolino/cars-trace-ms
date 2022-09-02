package process

import (
	"cars-data/process/config"
	"cars-data/process/models"
	"cars-data/process/repository"
	"cars-data/process/services/reader"
	"cars-data/process/services/splitter"

	"sync"
)

func Start() {
	var wg sync.WaitGroup

	data := reader.GetCsvData()

	slice01, slice02, slice03, slice04 := splitter.SplitCsvData(data)

	config := config.GetConfig()
	repository := &repository.Repository{}
	repository.InitializeDatabase(config)

	process(repository, &wg, slice01, slice02, slice03, slice04)

	wg.Wait()
}

func process(repo *repository.Repository, wg *sync.WaitGroup, slices ...[]models.CarEntity) {
	for _, slice := range slices {
		wg.Add(1)
		go repo.SaveCsvData(wg, slice)
	}
}
