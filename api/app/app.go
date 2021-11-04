package app

import (
	"github.com/gmvbr/go-rest/api/routes"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	api := app.Group("/public")
	routes.AuthRoute(api)
	return app
}
