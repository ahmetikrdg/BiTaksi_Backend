package controllers_Package_Test

import (
	"BiTaksi_Backend_Matching/controllers"
	"BiTaksi_Backend_Matching/models"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetByKmWithDriverLoction(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "Posts location information and returns 200 status codes",
			route:        "/match",
			expectedCode: 401,
		},
	}
	app := fiber.New()
	app.Post("/match", controllers.FindDriverWithLocation)
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

func TestGetByKmWithDriverLoction1(t *testing.T) {
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
	app.Post("/match", controllers.FindDriverWithLocation)
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
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByKmWithDriverLoction_Wrong_Model_Return_400(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "Posts location information and returns 200 status codes",
			route:        "/match",
			expectedCode: 404,
		},
	}
	app := fiber.New()
	app.Post("/match", controllers.FindDriverWithLocation)
	model := models.MatchingRequest{
		Type:        "Point",
		Coordinates: []float64{40.00009384, 20.200009532},
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
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestGetByKmWithDriverLoction21(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "Posts location information and returns 200 status codes",
			route:        "/match",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/match", controllers.FindDriverWithLocation)
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
		req.Header.Set("Authorization", "Bearer eyJhsadasdzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
