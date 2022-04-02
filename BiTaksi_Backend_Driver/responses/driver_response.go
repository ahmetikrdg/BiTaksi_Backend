package responses

import "github.com/gofiber/fiber/v2"

type DriverResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
