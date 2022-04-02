package routes

import (
	"BiTaksi_Backend_Matching/controllers"
	"github.com/gofiber/fiber/v2"
)

func MatchRoute(app *fiber.App) {
	app.Post("/match", controllers.FindDriverWithLocation)
}
