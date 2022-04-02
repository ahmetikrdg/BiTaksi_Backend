package controllers

import (
	models "BiTaksi_Backend_Matching/models"
	"BiTaksi_Backend_Matching/responses"
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"strings"
)

var validate = validator.New()

// FindDriverWithLocation godoc
// @Summary      Find The Driver
// @Description  The endpoint that allows searching with a GeoJSON point to find a driver if it matches the given criteria. Otherwise, the service should respond with a 404 - Not Found
// @Tags         Matching
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Authentication header: (If you want to use jwt, you have to write a bearer space token. Example Bearer eyJhbGciO)"
// @Param        match  body      models.MatchingRequest  true  "Rider data"
// @Success      200  {object}   responses.RiderResponse
// @Failure      400  {object}   responses.RiderResponse
// @Failure      404  {object}   responses.RiderResponse
// @Failure      401  {object}   responses.RiderResponse
// @Router       /match [post]
func FindDriverWithLocation(c *fiber.Ctx) error {
	var match models.MatchingRequest
	checkAuthStatus := false
	if err := c.BodyParser(&match); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	requestToken := c.GetReqHeaders()["Authorization"]
	if requestToken == "" || requestToken == "Bearer" {
		return c.Status(http.StatusUnauthorized).JSON(responses.RiderResponse{Status: http.StatusUnauthorized, Message: "error", Data: &fiber.Map{"data": "JWT Token not found"}})
	}

	newToken := strings.Fields(requestToken)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(newToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<true>"), nil
	})

	for key, val := range claims {
		if key == "name" && val == "Ahmet" {
			checkAuthStatus = true
		} else {
			checkAuthStatus = false
		}
	}

	if checkAuthStatus == false {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"data": "JUnauthorazation Data"}})
	}

	if validationErr := validate.Struct(&match); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	model := models.MatchingRequest{
		Id:          primitive.NewObjectID(),
		Type:        match.Type,
		Coordinates: match.Coordinates,
	}

	locationJson, err := json.Marshal(model)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/match", bytes.NewBuffer(locationJson))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	req.Header.Set("ApiKey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	body, err := ioutil.ReadAll(resp.Body)

	var result responses.RiderResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	if result.Status == 404 {
		return c.Status(http.StatusNotFound).JSON(responses.RiderResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"FATAL": "404"}})
	}
	return c.Status(http.StatusOK).JSON(responses.RiderResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Distance": result.Data}})
}
