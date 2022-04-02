package coordinatEvaluation

import (
	"BiTaksi_Backend_Driver/models"
)

func CoordinatControl(crd models.Coordinat, drivers []models.Coordinat) float64 {

	var km []float64

	for _, element := range drivers[1:] {
		h := Haversine(crd, element.Latitude, element.Longtitude)
		km = append(km, h)
	}

	calculated := KmCalculation(km)
	return calculated
}
