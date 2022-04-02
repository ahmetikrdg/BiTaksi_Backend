package haversine

import (
	"BiTaksi_Backend_Driver/configs"
	"BiTaksi_Backend_Driver/models"
	"BiTaksi_Backend_Driver/tools/zap_logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "driver")

func CoordinatControls(crd models.Coordinat) float64 {
	zap_logger.ServerInfoWithInfoLog("Starting the CoordinatControls process")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.DriverLocation
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.DriverLocation
		if err = results.Decode(&singleUser); err != nil {
			fmt.Println(err)
		}
		users = append(users, singleUser)
	}

	var count int = 0
	var km []float64
	crdLat := fmt.Sprintf("%.3f", crd.Latitude)
	crdLon := fmt.Sprintf("%.3f", crd.Longtitude)

	for _, element := range users[1:] {
		elementLat := fmt.Sprintf("%.3f", element.Location.Coordinates[0])
		elementLon := fmt.Sprintf("%.3f", element.Location.Coordinates[1])

		if elementLat == crdLat || elementLon == crdLon { //başlayan konumları al
			h := Haversine(crd, element.Location.Coordinates[0], element.Location.Coordinates[1])
			km = append(km, h)
			count++
		}
	}
	zap_logger.ServerInfoWithInfoLog("the data has arrived")

	fmt.Println("Data: ", count)

	if count == 0 {
		return 0
	}

	var smallest float64 = km[0]
	for _, num := range km {
		if num < smallest {
			smallest = num
		}
	}

	zap_logger.ServerInfoWithInfoLog("CoordinatControls successful")
	return smallest
}
