package main

import (
	"log"
	"os"

	"github.com/davidfunk13/overwatch-companion/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var CorsConfig = cors.Config{
		Next:             nil,
		AllowOrigins:     "https://overwatch-companion.netlify.app",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}

	app := fiber.New()

	app.Use(cors.New(CorsConfig))

	port := os.Getenv("PORT")

	//if is dev
	if port == "" {
		port = "3001"
	}

	api.RouteHandlers(app)

	log.Fatal(app.Listen(":" + port))
}
