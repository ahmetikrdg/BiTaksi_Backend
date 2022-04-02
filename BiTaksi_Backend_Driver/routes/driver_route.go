package routes

import (
	"BiTaksi_Backend_Driver/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/bulkOperations", controllers.CreateAndUpdate)
	app.Get("/driver/:driverId", controllers.GetADriver)
	app.Delete("/driver/:driverId", controllers.DeleteADriver)
	app.Get("/drivers", controllers.GetAllDrivers)
	app.Post("/match", controllers.GetByLocationBase)
	app.Post("/matchTwo", controllers.GetByLocations)

}
