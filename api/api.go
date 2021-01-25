package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ConfigureRouter : Sets up all routes for app.
func ConfigureRouter(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		fmt.Print("Testing")
		return c.SendString("Hello, World!")
	})
}
