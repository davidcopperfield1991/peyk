package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RunFiber() {
	app := fiber.New()

	app.Get("/message", func(c *fiber.Ctx) error {
		return c.SendString("message")
	})

	log.Fatal(app.Listen(":6060"))
}
