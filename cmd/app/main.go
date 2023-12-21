package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("../../tmp", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		})

	app.Static("/", "../../tmp")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("main", fiber.Map{})
	})

	app.Listen(":3000")
}
