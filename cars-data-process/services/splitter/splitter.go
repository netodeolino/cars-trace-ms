package splitter

import (
	"github.com/netodeolino/cars-trace-ms/cars-data-core/models"
)

func SplitCsvData(data []models.CarModel) (sl1, sl2, sl3, sl4 []models.CarModel) {
	length := len(data)

	var fisrtPart int = (length / 4)
	var secondPart int = (fisrtPart * 2)
	var thirdPart int = (fisrtPart * 3)

	slice1 := data[:fisrtPart]
	slice2 := data[fisrtPart:secondPart]
	slice3 := data[secondPart:thirdPart]
	slice4 := data[thirdPart:]

	return slice1, slice2, slice3, slice4
}
