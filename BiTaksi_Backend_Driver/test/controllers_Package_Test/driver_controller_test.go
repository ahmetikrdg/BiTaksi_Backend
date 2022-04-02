package controllers_Package_Test

import (
	"BiTaksi_Backend_Driver/controllers"
	"BiTaksi_Backend_Driver/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetByLocationBase_Return_200_OK(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "Posts location information and returns 200 status codes",
			route:        "/match",
			expectedCode: 200,
		},
	}
	app := fiber.New()
	app.Post("/match", controllers.GetByLocationBase)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{41.81179384, 29.22299532},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("ApiKey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocationBase_Wrong_Model(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		// First test case
		{
			description:  "Posts wrong model and gives 500 status code in return",
			route:        "http://localhost:8000/match",
			expectedCode: 500,
		},
	}
	app := fiber.New()
	app.Post("/match", controllers.GetByLocationBase)

	locationJson, err := json.Marshal("model")
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("ApiKey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocationBase_Empty_Model_400(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Post sends blank model and receives 400 status codes",
			route:        "http://localhost:8000/match",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/match", controllers.GetByLocationBase)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{0, 0},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("ApiKey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestDeleteADriver(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		// First test case
		{
			description:  "Posts id info and drive is deleted",
			route:        "/driver/624865a5ac65ce9dd4e2f913",
			expectedCode: 200,
		},
		{
			description:  "Non-existent page writes and gets error",
			route:        "/not-found",
			expectedCode: 404,
		},
		{
			description:  "Gets routing wrong and gets 404 error",
			route:        "/driver/62460f843b580a99a3fda8a6/2",
			expectedCode: 404,
		},
		{
			description:  "Posts id less than 23 elements and gets 400 error",
			route:        "/driver/6246dfdfdf3b9a3fda8d8",
			expectedCode: 400,
		},
	}

	// Define Fiber app.
	app := fiber.New()
	app.Delete("/driver/:driverId", controllers.DeleteADriver)

	for _, test := range tests {

		req := httptest.NewRequest("DELETE", test.route, nil)
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocations_Return_200_404(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "post HTTP status 200",
			route:        "/matchTwo",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}
	app := fiber.New()
	app.Post("/matchTwo", controllers.GetByLocations)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{40.11628248, 29.09388969},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocations_Wrong_Model_Return_400(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts location information and returns 200 status codes",
			route:        "/matchTwo",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/matchTwo", controllers.GetByLocations)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{0, 0},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocations_Wrong_Data_404(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts wrong model and gives 500 status code in return",
			route:        "/matchTwo",
			expectedCode: 404,
		},
	}
	app := fiber.New()
	app.Post("/matchTwo", controllers.GetByLocations)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{40.0000988, 29.0007119},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocations_Return_422(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "",
			route:        "/matchTwo",
			expectedCode: 422,
		},
	}
	app := fiber.New()
	app.Post("/matchTwo", controllers.GetByLocations)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{41.08678928, 29.09087169},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByLocations_Return_400(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts wrong model and gets 400 error\n",
			route:        "/matchTwo",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/matchTwo", controllers.GetByLocations)
	model := models.MatchingRequest{
		Coordinates: []float64{41.08678928, 29.09087169},
	}
	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateAndUpdate(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts blank model and receives 400 status codes",
			route:        "/bulkOperations",
			expectedCode: 400,
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}
	app := fiber.New()
	app.Post("/bulkOperations", controllers.CreateAndUpdate)

	model := models.DriverLocation{}

	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateAndUpdate_StatusCode_400(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts wrong model and receives 400 status codes",
			route:        "/bulkOperations",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/bulkOperations", controllers.CreateAndUpdate)
	rider := models.Location{
		Type:        "denemememe",
		Coordinates: []float64{05464565565645.54, 27789789787897989.09087169},
	}
	model := models.DriverLocation{Location: rider}

	locationJson, err := json.Marshal(model)
	if err != nil {
		assert.Error(t, err)
	}

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateAndUpdate_StatusCode_201_One_Create(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts correct model and status code 201\n",
			route:        "/bulkOperations",
			expectedCode: 201,
		},
	}

	rider := models.Location{
		Type:        "Point",
		Coordinates: []float64{41.11678328, 29.09087169},
	}

	model := models.DriverLocation{Location: rider}

	var drivers []models.DriverLocation
	drivers = append(drivers, model)

	locationJson, err := json.Marshal(drivers)
	if err != nil {
		assert.Error(t, err)
	}

	app := fiber.New()
	app.Post("/bulkOperations", controllers.CreateAndUpdate)

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateAndUpdate_StatusCode_201_Multiple_Create(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts multiple data and receives status code 201",
			route:        "/bulkOperations",
			expectedCode: 201,
		},
	}

	rider := models.Location{
		Type:        "Point",
		Coordinates: []float64{42.11678328, 25.09087169},
	}
	riderOne := models.Location{
		Type:        "Point",
		Coordinates: []float64{41.11628328, 26.09085169},
	}
	riderTwo := models.Location{
		Type:        "Point",
		Coordinates: []float64{45.11378528, 28.09087169},
	}

	model := models.DriverLocation{Location: rider}
	modelOne := models.DriverLocation{Location: riderOne}
	modelTwo := models.DriverLocation{Location: riderTwo}

	var drivers []models.DriverLocation
	drivers = append(drivers, model)
	drivers = append(drivers, modelOne)
	drivers = append(drivers, modelTwo)

	locationJson, err := json.Marshal(drivers)
	if err != nil {
		assert.Error(t, err)
	}

	app := fiber.New()
	app.Post("/bulkOperations", controllers.CreateAndUpdate)

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateAndUpdate_StatusCode_201_Create_and_Update(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts multiple data and receives status code 201",
			route:        "/bulkOperations",
			expectedCode: 201,
		},
	}

	rider := models.Location{
		Type:        "Point",
		Coordinates: []float64{40.11628328, 29.09387169},
	}

	riderOne := models.Location{
		Type:        "Point",
		Coordinates: []float64{40.11628328, 29.09387169},
	}

	id, _ := primitive.ObjectIDFromHex(fmt.Sprintf("%s", "624857fef4e6afdcf8fe9f02"))
	model := models.DriverLocation{Id: id, Location: rider}
	modelOne := models.DriverLocation{Location: riderOne}

	var drivers []models.DriverLocation
	drivers = append(drivers, model)
	drivers = append(drivers, modelOne)

	locationJson, err := json.Marshal(drivers)
	if err != nil {
		assert.Error(t, err)
	}

	app := fiber.New()
	app.Post("/bulkOperations", controllers.CreateAndUpdate)

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestCreateAndUpdate_StatusCode_500_Create(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Posts blank model and receives 500 status codes",
			route:        "/bulkOperations",
			expectedCode: 500,
		},
	}

	rider := models.Location{}

	model := models.DriverLocation{Location: rider}

	var drivers []models.DriverLocation
	drivers = append(drivers, model)

	locationJson, err := json.Marshal(drivers)
	if err != nil {
		assert.Error(t, err)
	}

	app := fiber.New()
	app.Post("/bulkOperations", controllers.CreateAndUpdate)

	// Iterate through test single test cases
	for _, test := range tests {
		req, err := http.NewRequest("POST", test.route, bytes.NewBuffer(locationJson))
		if err != nil {
			assert.Error(t, err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetADriver_Return_StatusCode(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Gives wrong driver ID and gets 404 status code",
			route:        "/driver/62460fzzzzb580a99a3fda8a2",
			expectedCode: 404,
		},
		{
			description:  "Gives wrong driver ID and gets 404 status code",
			route:        "/driver/62460f843b580a99a3fda8a4/1",
			expectedCode: 404,
		},
		//  test case
		{
			description:  "Posts driver id and receives 200 status codes",
			route:        "/driver/624865a5ac65ce9dd4e2f8fc",
			expectedCode: 200,
		},
	}

	// Define Fiber app.
	app := fiber.New()
	app.Get("/driver/:driverId", controllers.GetADriver)

	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}
}

func TestGetADriver_Return_StatusCode_400_404(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Gives wrong driver ID and gets 404 status code",
			route:        "/driver/62460f843b580a99a3fda8",
			expectedCode: 400,
		},
		// First test case
		{
			description:  "Gives wrong driver ID and gets 404 status code",
			route:        "/driver/34160f843b580b11a3fda8d2",
			expectedCode: 404,
		},
	}

	// Define Fiber app.
	app := fiber.New()
	app.Get("/driver/:driverId", controllers.GetADriver)

	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}
}

func TestGetAllDrivers(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "Fetch all data",
			route:        "/drivers",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "Gets HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
		// Third test case
		{
			description:  "Gets bad page and gets 404 status code",
			route:        "/drivers/asdasd",
			expectedCode: 404,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/drivers", controllers.GetAllDrivers)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
