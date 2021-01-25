package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	env := os.Getenv("APP_ENV")
	fmt.Println(reflect.TypeOf(port))
	fmt.Println(reflect.TypeOf(env))
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Print("Testing")
		return c.SendString("Hello, World!")
	})
	app.Listen(port)
}
