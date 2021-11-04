package routes

import "github.com/gofiber/fiber/v2"

func AuthRoute(app fiber.Router) {
	app.Post("/auth", auth())
}

func auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "Hello world!",
		})
	}
}
