package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Print("Testing")
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}
