package main

import (
	"log"
	"os"

	"github.com/davidfunk13/overwatch-companion/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://overwatch-companion.netlify.app/, http://localhost:3000, http://localhost:3001",
	}))

	port := os.Getenv("PORT")

	//if is dev
	if port == "" {
		port = "3001"
	}

	api.RouteHandlers(app)

	log.Fatal(app.Listen(":" + port))
}
