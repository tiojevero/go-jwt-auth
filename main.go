package main

import (
	"github.com/gofiber/fiber/v2"
)

const port string = ":5000"

type SignupRequest struct {
	Name string
	Email string
	Password string
}

func main() {
	app := fiber.New()

	_, err := createDBEngine()
	if err != nil {
		panic(err)
	}

	app.Post("/signup", func(c *fiber.Ctx) error {
		req := new(SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.Name == "" || req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid signup request")
		}

		

		return nil
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return nil
	})

	app.Post("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path":"private"})
	})

	app.Post("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path":"public"})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	} )

	app.Listen(port)
}


