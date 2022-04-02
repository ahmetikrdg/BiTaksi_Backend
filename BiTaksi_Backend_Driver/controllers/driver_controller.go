package controllers

import (
	"BiTaksi_Backend_Driver/configs"
	"BiTaksi_Backend_Driver/models"
	"BiTaksi_Backend_Driver/repository/mongodb"
	"BiTaksi_Backend_Driver/responses"
	coordinatEvalution "BiTaksi_Backend_Driver/tools/coordinatEvaluation"
	"BiTaksi_Backend_Driver/tools/errors"
	"BiTaksi_Backend_Driver/tools/haversine"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var driverCollection *mongo.Collection = configs.GetCollection(configs.DB, "driver")
var validate = validator.New()

//CreateAndUpdate godoc
//@Summary      Bulk Create and Update
//@Description  An endpoint for creating a driver location. It would support batch operations to handle the bulk update.
//@Tags         Driver
//@Accept       json
//@Produce      json
//@Param        drivers  body     []models.DriverLocation  true  "Driver data"
//@Success      200  {object}   responses.DriverResponse
//@Failure      400  {object}   responses.DriverResponse
//@Failure      404  {object}   responses.DriverResponse
//@Failure      401  {object}   responses.DriverResponse
//@Router       /bulkOperations [post]
func CreateAndUpdate(c *fiber.Ctx) error {
	var drivers []models.DriverLocation

	var insertCount []*mongo.InsertOneResult
	var updateCount int64

	if err := c.BodyParser(&drivers); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	for _, driver := range drivers {
		if driver.Id.IsZero() == true { //create
			newDriver := models.DriverLocation{
				Id:       primitive.NewObjectID(),
				Location: driver.Location,
			}

			result, err := mongodb.Insert(newDriver)

			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"message": err.Error()}})
			}
			insertCount = append(insertCount, result)
		}

		_, count, err := mongodb.Update(driver)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"message": err.Error()}})
		}

		if count == 1 {
			updateCount += count
		}

	}

	if updateCount != 0 && len(insertCount) != 0 {
		return c.Status(http.StatusCreated).JSON(responses.DriverResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "Add and Update process complete."}})

	}

	if updateCount != 0 {
		return c.Status(http.StatusCreated).JSON(responses.DriverResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"Affected Update Data": updateCount}})

	}

	if len(insertCount) != 0 {
		return c.Status(http.StatusCreated).JSON(responses.DriverResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"Affected Insert Data": len(insertCount)}})
	}

	return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"message": "FATAL"}})
}

//GetAllDrivers godoc
//@Summary      Fetch All Drivers
//@Description  This endpoint fetches all the driver data
//@Tags         Driver
//@Accept       json
//@Produce      json
//@Success      200  {object}   responses.DriverResponse
//@Failure      400  {object}   responses.DriverResponse
//@Failure      404  {object}   responses.DriverResponse
//@Failure      401  {object}   responses.DriverResponse
//@Router       /drivers [get]
func GetAllDrivers(c *fiber.Ctx) error {
	results, err := mongodb.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"message": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(
		responses.DriverResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": results}},
	)
}

//GetADriver godoc
//@Summary      Search for a Driver
//@Description  An endpoint finds a driver by the given id
//@ID get-string-by-int
//@Tags         Driver
//@Accept       json
//@Produce      json
//@Param        driverId path     string true "Driver ID"
//@Success      200  {object}   responses.DriverResponse
//@Failure      400  {object}   responses.DriverResponse
//@Failure      404  {object}   responses.DriverResponse
//@Failure      401  {object}   responses.DriverResponse
//@Router       /driver/{driverId} [get]
func GetADriver(c *fiber.Ctx) error {
	driverId := c.Params("driverId")

	if len(driverId) < 23 {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"message": "The request structure is incorrect, please enter a correct id"}})
	}

	result, err1, err2 := mongodb.GetById(driverId)

	if err1 != nil && err2 != nil || err1 == mongo.ErrNoDocuments && err2 == mongo.ErrNoDocuments {
		return c.Status(http.StatusNotFound).JSON(responses.DriverResponse{Status: http.StatusNotFound, Message: "No Content", Data: &fiber.Map{"Information": "There are no taxis near you"}})
	}

	return c.Status(http.StatusOK).JSON(responses.DriverResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})

}

//GetByLocations godoc
//@Summary      Matching a Driver Location Endpoint
//@Description  This endpoint searches with a different path and algorithm. Jwt doesn't want to.
//@Tags         Driver
//@Accept       json
//@Produce      json
//@Param        driver  body     models.RiderRequest  true  "Driver data"
//@Success      200  {object}   responses.DriverResponse
//@Failure      400  {object}   responses.DriverResponse
//@Failure      404  {object}   responses.DriverResponse
//@Failure      401  {object}   responses.DriverResponse
//@Router       /matchTwo [post]
func GetByLocations(c *fiber.Ctx) error {
	var driver models.MatchingRequest
	if err := c.BodyParser(&driver); err != nil {
		return err
	}

	if validationErr := validate.Struct(&driver); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"message": "You must enter numbers other than zero"}})
	}

	inComing := models.Coordinat{
		Latitude:   driver.Coordinates[0],
		Longtitude: driver.Coordinates[1],
	}

	if inComing.Latitude == 0 && inComing.Longtitude == 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"message": "You must enter numbers other than zero"}})
	}

	a := haversine.CoordinatControls(inComing)
	if a == 0 {
		return c.Status(http.StatusNotFound).JSON(responses.DriverResponse{Status: http.StatusNotFound, Message: "No Content", Data: &fiber.Map{"Information": "There are no taxis near you"}}) //burayı yakında diye değiştir aynı konumu verirse 0km gelir
	}

	return c.Status(http.StatusOK).JSON(responses.DriverResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"KM": a}})

}

//GetByLocationBase godoc
//@Summary      Matching a Driver Location Endpoint
//@Description  An endpoint for matches the nearest driver with the passenger
//@Tags         Driver
//@Accept       json
//@Produce      json
//@Param        driver  body     models.RiderRequest  true  "Driver data"
//@Success      200  {object}   responses.DriverResponse
//@Failure      400  {object}   responses.DriverResponse
//@Failure      404  {object}   responses.DriverResponse
//@Failure      401  {object}   responses.DriverResponse
//@Router       /match [post]
func GetByLocationBase(c *fiber.Ctx) error {
	checkAuthStatus := false
	var driver models.MatchingRequest
	secretkeyFromMatching := c.GetReqHeaders()["Apikey"]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(secretkeyFromMatching, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("ApiKey"), nil
	})

	for key, val := range claims {
		if key == "name" && val == "Ahmet" {
			checkAuthStatus = true
		} else {
			checkAuthStatus = false
		}
	}

	if checkAuthStatus == false {
		return c.Status(http.StatusUnauthorized).JSON(responses.DriverResponse{Status: http.StatusUnauthorized, Message: "error", Data: &fiber.Map{"message": "Unauthorazation Data"}})
	}
	//------

	if err := c.BodyParser(&driver); err != nil {
		return err
	}
	//defer cancel()
	if validationErr := validate.Struct(&driver); validationErr != nil {
		return err
	}

	inComingCoordinat := models.Coordinat{
		Latitude:   driver.Coordinates[0],
		Longtitude: driver.Coordinates[1],
	}

	if inComingCoordinat.Latitude == 0 && inComingCoordinat.Longtitude == 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"message": "You must enter numbers other than zero"}})
	}

	coordinatDatas, err := mongodb.NearSphere(inComingCoordinat)
	if coordinatDatas == nil {
		return c.Status(http.StatusNotFound).JSON(responses.DriverResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"FATAL": "404"}})
	}

	errors.StandartErrorWithErrorLog(err, nil)
	distance := coordinatEvalution.CoordinatControl(inComingCoordinat, coordinatDatas)

	if distance < 0 {
		return c.Status(http.StatusNotFound).JSON(responses.DriverResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"FATAL": "404"}})

	}

	return c.Status(http.StatusOK).JSON(responses.DriverResponse{Data: &fiber.Map{"KM": distance}})
}

//DeleteADriver godoc
//@Summary      Delete a Driver Location
//@Description  This endpoint deletes a selected data by id
//@ID get-string-by-int
//@Tags         Driver
//@Accept       json
//@Produce      json
//@Param        driverId path     string true "Driver ID"
//@Success      200  {object}   responses.DriverResponse
//@Failure      400  {object}   responses.DriverResponse
//@Failure      404  {object}   responses.DriverResponse
//@Failure      401  {object}   responses.DriverResponse
//@Router       /driver/{driverId} [delete]
func DeleteADriver(c *fiber.Ctx) error {
	driverId := c.Params("driverId")
	if len(driverId) < 23 {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"message": "The request structure is incorrect, please enter a correct id"}})
	}

	result, err := mongodb.Delete(driverId)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"message": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.DriverResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"Fatal": "404"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.DriverResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Driver successfully deleted!"}},
	)
}
