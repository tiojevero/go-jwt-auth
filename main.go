package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const port string = ":5000"

type SignupRequest struct {
	Name string
	Email string
	Password string
}

type LoginRequest struct {
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
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); 

		if err != nil {
			return err
		}

		user := &User {
			Name : req.Name,
			Email : req.Email,
			Password : string(hash),
		}
		
		checkUser := new(User)
		has, err := engine.Where("email = ?", req.Email).Desc("id").Get(checkUser)
		if err != nil {
			return err
		}
		if has {
			return fiber.NewError(fiber.StatusBadRequest, "Email Already Registered")
		} 

		_, err = engine.Insert(user) 
		if err != nil {
			return err
		}


		// CREATE JWT TOKEN
		token, exp, err := createJWTToken(*user)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		req := new(LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid Login Credentials")
		}


		// FIND USER IN DATABASE
		user := new(User)
		has, err := engine.Where("email = ?", req.Email).Desc("id").Get(user)
		fmt.Println(has)
		if err != nil {
			return err
		}

		if !has {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid Login Credentials")
		}


		// CREATE JWT TOKEN
		token, exp, err := createJWTToken(*user)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
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


func createJWTToken(user User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id 
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}