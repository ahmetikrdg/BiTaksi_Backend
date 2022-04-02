package haversine

import (
	"BiTaksi_Backend_Driver/models"
	"math"
)

func Haversine(cr models.Coordinat, lat1 float64, lon1 float64) float64 {
	var dLat float64 = (math.Pi / 180) * (cr.Latitude - lat1)
	var dLon float64 = (math.Pi / 180) * (cr.Longtitude - lon1)

	lat1 = (math.Pi / 180) * (lat1)
	cr.Latitude = (math.Pi / 180) * (cr.Latitude)

	var a float64 = math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*math.Cos(lat1)*math.Cos(cr.Latitude)
	var rad float64 = 6371
	var c float64 = 2 * math.Asin(math.Sqrt(a))
	return rad * c
}
