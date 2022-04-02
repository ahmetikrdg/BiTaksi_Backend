package main

import (
	_ "BiTaksi_Backend_Matching/docs"
	"BiTaksi_Backend_Matching/routes"
	Swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title           Matching API
// @version         1.0
// @description     Find the nearest driver to rider.

// @contact.name   Ahmet Karadağ
// @contact.url    http://github.com/ahmetikrdg
// @contact.email  ahmetikrdg@outlook.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /

// @in                          header
// @name                        Authorization
func main() {
	app := fiber.New()
	app.Get("/swagger/*", Swagger.HandlerDefault)
	routes.MatchRoute(app)
	app.Listen(":5000")
}
