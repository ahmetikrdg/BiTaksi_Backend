package main

import (
	"BiTaksi_Backend_Driver/configs"
	_ "BiTaksi_Backend_Driver/docs"
	"BiTaksi_Backend_Driver/routes"
	Swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

//@title           Driver Location API
//@version 1.0
//@description Find the nearest driver to rider.

//@contact.name   API
//@contact.url    http://github.com/ahmetikrdg
//@contact.email  ahmetikrdg@outlook.com

//@license.name  Apache 2.0
//@license.url   http://apache.org/licenses/LICENSE-2.0.html

//@host localhost:8000
//@BasePath /
func main() {
	app := fiber.New()
	app.Get("/swagger/*", Swagger.HandlerDefault)
	configs.CreateLocationData()
	configs.ConnectDB()
	routes.UserRoute(app)
	app.Listen(":8000")
}
