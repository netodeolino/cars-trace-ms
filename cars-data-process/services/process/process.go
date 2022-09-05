package process

import (
	"log"

	"github.com/netodeolino/cars-trace-ms/cars-data-core/config"
	"github.com/netodeolino/cars-trace-ms/cars-data-core/models"
	"github.com/netodeolino/cars-trace-ms/cars-data-core/repository"
	"github.com/netodeolino/cars-trace-ms/cars-data-process/services/reader"
	"github.com/netodeolino/cars-trace-ms/cars-data-process/services/splitter"

	"sync"
)

func Start() {
	var wg sync.WaitGroup

	data := reader.GetCsvData()

	slice01, slice02, slice03, slice04 := splitter.SplitCsvData(data)

	config := config.GetConfig()
	repository, err := repository.NewRepository(config)

	if err != nil {
		log.Fatal(err)
	}

	process(repository, &wg, slice01, slice02, slice03, slice04)

	wg.Wait()
}

func process(repo *repository.Repository, wg *sync.WaitGroup, slices ...[]models.CarModel) {
	for _, slice := range slices {
		wg.Add(1)
		go repo.SaveCsvData(wg, slice)
	}
}
