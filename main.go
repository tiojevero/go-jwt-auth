package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const port string = ":5000"

type SignupRequest struct {
	Name string
	Email string
	Password string
}

func main() {
	app := fiber.New()

	engine, err := createDBEngine()
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

		// SAVE THIS INFO IN TO THE DATABASE
		hash, err := bcrypt.GenerateFromPassword([]byte{req.Password}, bcrypt.DefaultCost); 

		if err != nil {
			return err
		}

		user := &User {
			Name : req.Name,
			Email : req.Email,
			Password : string(hash),
		}

		_, err = engine.Insert(user) 
		if err != nil {
			return err
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


