package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

func StartServer() {
	engine := html.New("../../tmp", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		})

	app.Static("/", "../../tmp")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("main", fiber.Map{})
	})

	log.Trace("test")

	app.Listen(":3000")
}
