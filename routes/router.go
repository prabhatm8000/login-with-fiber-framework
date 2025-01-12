package routes

import (
	"example.com/login/routes/authRoutes"
	"github.com/gofiber/fiber/v2"
)

func SetupAPIRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is up and running!")
	})
	authRoutes.AddLoginRoute(api)
}
