package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Text   string `json:"text"`
	Sender string `json:"sender"`
	To     string `json:"to"`
}

func sendMessage(c *fiber.Ctx) error {
	data := Message{}
	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Text)
	fmt.Println(data.Sender)
	fmt.Println(data.To)

	return c.JSON(data)
}

func RunFiber() {
	app := fiber.New()

	app.Get("/message", func(c *fiber.Ctx) error {
		return c.SendString("message")
	})

	app.Post("/send/:message/:sender/:to", sendMessage)

	log.Fatal(app.Listen(":6060"))

}
