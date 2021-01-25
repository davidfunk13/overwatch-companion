package main

import (
	"log"
	"os"

	"github.com/davidfunk13/overwatch-companion/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	//if is dev
	if port == "" {
		port = "3001"
	}

	api.ConfigureRouter(app)

	log.Fatal(app.Listen(":" + port))
}
