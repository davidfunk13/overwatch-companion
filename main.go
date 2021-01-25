package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Println("Port:", port)

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Print("Testing")
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + port)
}
