package mongodb

import (
	"BiTaksi_Backend_Driver/configs"
	"BiTaksi_Backend_Driver/models"
	"BiTaksi_Backend_Driver/tools/errors"
	"BiTaksi_Backend_Driver/tools/zap_logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var driverCollection *mongo.Collection = configs.GetCollection(configs.DB, "driver")

func Insert(driver models.DriverLocation) (*mongo.InsertOneResult, error) {
	zap_logger.ServerInfoWithInfoLog("Starting the Insert process")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := driverCollection.InsertOne(ctx, driver)

	errors.ServerErrorWithErrorLog(err)
	zap_logger.ServerInfoWithInfoLog("Insert successful")
	return result, err
}

func Update(driver models.DriverLocation) (*mongo.UpdateResult, int64, error) {
	zap_logger.ServerInfoWithInfoLog("Starting the Update process")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(driver.Id.Hex())
	errors.ServerErrorWithErrorLog(err)
	update := bson.M{"location.type": driver.Location.Type, "location.coordinates": driver.Location.Coordinates}

	result, err := driverCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	errors.ServerErrorWithErrorLog(err)
	modified := result.ModifiedCount

	zap_logger.ServerInfoWithInfoLog("Update successful")
	return result, modified, err
}

func GetAll() ([]models.DriverLocation, error) {
	zap_logger.ServerInfoWithInfoLog("Starting the GetAll process")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var drivers []models.DriverLocation
	defer cancel()

	results, err := driverCollection.Find(ctx, bson.M{})

	errors.ServerErrorWithErrorLog(err)

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleDriver models.DriverLocation
		if err = results.Decode(&singleDriver); err != nil {
			return nil, err
		}

		drivers = append(drivers, singleDriver)
	}

	zap_logger.ServerInfoWithInfoLog("GetAll successful")
	return drivers, nil
}

func GetById(id string) (models.DriverLocation, error, error) {
	zap_logger.ServerInfoWithInfoLog("Starting the GetById process")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var driver models.DriverLocation
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err1 := driverCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&driver)
	err2 := driverCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&driver)

	zap_logger.ServerInfoWithInfoLog("GetById successful")
	return driver, err1, err2
}

func NearSphere(coordinat models.Coordinat) ([]models.Coordinat, error) {
	zap_logger.ServerInfoWithInfoLog("Starting the NearSphere process")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := driverCollection.Find(ctx, bson.M{"location": bson.M{
		"$nearSphere": bson.M{
			"$geometry": bson.M{
				"type":        "Point",
				"coordinates": []float64{coordinat.Latitude, coordinat.Longtitude},
			},
			"$maxDistance": 10 * 1000,
		},
	}})
	errors.ServerErrorWithErrorLog(err)

	var coordinatDatas []models.Coordinat

	defer results.Close(ctx)
	for results.Next(ctx) {
		var driverLocation models.DriverLocation
		if err = results.Decode(&driverLocation); err != nil {
			fmt.Println(err)
		}

		nearCoordinat := models.Coordinat{
			Latitude:   driverLocation.Location.Coordinates[0],
			Longtitude: driverLocation.Location.Coordinates[1],
		}
		coordinatDatas = append(coordinatDatas, nearCoordinat)
	}

	if coordinatDatas == nil {
		return nil, err
	}
	zap_logger.ServerInfoWithInfoLog("NearSphere successful")
	return coordinatDatas, nil
}

func Delete(id string) (*mongo.DeleteResult, error) {
	zap_logger.ServerInfoWithInfoLog("Starting the deletion process")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(id)
	errors.ServerErrorWithErrorLog(err)
	result, err := driverCollection.DeleteOne(ctx, bson.M{"id": objId})
	errors.ServerErrorWithErrorLog(err)
	zap_logger.ServerInfoWithInfoLog("Deletion successful")
	return result, err
}
