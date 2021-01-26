package api

import (
	"github.com/davidfunk13/overwatch-companion/models"
	"github.com/gofiber/fiber/v2"
)

var sampleheroes = []models.Hero{
	{ID: 1, Name: "Junkrat", Role: 1},
	{ID: 2, Name: "Echo", Role: 1},
	{ID: 3, Name: "Wrecking Ball", Role: 0},
	{ID: 4, Name: "RoadHog", Role: 0},
	{ID: 5, Name: "Baptiste", Role: 2},
	{ID: 6, Name: "Zenyatta", Role: 2},
}

var sampleusers = []models.User{
	{ID: 1, Battletag: "buttkegels", Identifier: 1856},
	{ID: 2, Battletag: "buttkegels", Identifier: 1990},
	{ID: 3, Battletag: "NakedDave", Identifier: 11750},
}

// RouteHandlers : Sets up all routes for app.
func RouteHandlers(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		return c.JSON(sampleusers)
	})

	app.Get("/api/heroes", func(c *fiber.Ctx) error {
		return c.JSON(sampleheroes)
	})

}
