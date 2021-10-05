package main

import (
	"github.com/gofiber/fiber/v2"
)

const port string = ":5000"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	} )

	app.Listen(port)
}


