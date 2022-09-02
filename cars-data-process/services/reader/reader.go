package reader

import (
	"cars-data/process/models"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func GetCsvData() []models.CarEntity {
	file := openCsv()
	lines := readCsv(file)
	return mountCsvData(lines)
}

func openCsv() *os.File {
	file, err := os.Open("receiver/car-data.csv")

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func readCsv(file *os.File) [][]string {
	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return lines
}

func mountCsvData(lines [][]string) []models.CarEntity {
	carsData := []models.CarEntity{}
	for index, line := range lines {
		if index > 0 {
			latitude, _ := strconv.ParseFloat(line[1], 64)
			longitude, _ := strconv.ParseFloat(line[2], 64)
			stopped, _ := strconv.ParseBool(line[4])

			car := models.CarEntity{
				CarID:     line[0],
				Latitude:  latitude,
				Longitude: longitude,
				Operator:  line[3],
				Stopped:   stopped,
				LineId:    line[5],
			}

			carsData = append(carsData, car)
		}
	}

	return carsData
}
