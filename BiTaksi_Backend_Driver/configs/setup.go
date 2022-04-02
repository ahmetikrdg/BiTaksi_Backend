package configs

import (
	"BiTaksi_Backend_Driver/models"
	"BiTaksi_Backend_Driver/tools/convert"
	"BiTaksi_Backend_Driver/tools/errors"
	"BiTaksi_Backend_Driver/tools/zap_logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	zap_logger.ServerInfoWithInfoLog("ConnectDB the GetCollection process")
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	errors.ServerErrorWithErrorLog(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	errors.ServerErrorWithErrorLog(err)
	err = client.Ping(ctx, nil)
	errors.ServerErrorWithErrorLog(err)

	zap_logger.ServerInfoWithInfoLog("Connected to MongoDB")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	zap_logger.ServerInfoWithInfoLog("Starting the GetCollection process")
	collection := client.Database("BiTaksi").Collection(collectionName)
	zap_logger.ServerInfoWithInfoLog("GetCollection successful")
	return collection
}

func CreateLocationData() error {
	zap_logger.ServerInfoWithInfoLog("Starting the CreateLocationData process")
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	var driverCollection *mongo.Collection = GetCollection(DB, "driver")
	var isEmpty models.DriverLocation
	defer cancel()

	err := driverCollection.FindOne(ctx, bson.M{}).Decode(&isEmpty)
	errors.ServerErrorWithErrorLog(err)

	//--
	var onlyDriverLocationModel models.DriverLocation
	var onlyDriverLocation models.Location
	var DriverLocations []models.DriverLocation

	if isEmpty.Id.IsZero() == true {
		zap_logger.ServerInfoWithInfoLog("Collection Createing")
		locationData := convert.CsvToStruct()

		for _, element := range locationData {
			var location []float64

			location = append(location, element.Latitude)
			location = append(location, element.Longtitude)

			onlyDriverLocation.Type = "Point"
			onlyDriverLocation.Coordinates = location

			onlyDriverLocationModel.Id = primitive.NewObjectID()
			onlyDriverLocationModel.Location = onlyDriverLocation

			DriverLocations = append(DriverLocations, onlyDriverLocationModel)
		}

		var driverLocationData []interface{}

		for _, t := range DriverLocations[1:] {
			driverLocationData = append(driverLocationData, t)
		}
		_, err := driverCollection.InsertMany(ctx, driverLocationData)

		errors.ServerErrorWithErrorLog(err)

		Index := mongo.IndexModel{
			Keys: bson.D{
				{Key: "location", Value: "2dsphere"},
			},
		}

		_, err = driverCollection.Indexes().CreateOne(context.Background(), Index)
		errors.ServerErrorWithErrorLog(err)

		zap_logger.ServerInfoWithInfoLog("Data Ready")
		zap_logger.ServerInfoWithInfoLog("CreateLocationData successful")

		return nil
	}
	zap_logger.ServerInfoWithInfoLog("Data Already Exists")
	return nil
}
