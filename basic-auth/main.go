package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	// Setup Basic Auth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"test": "test",
		},
	}))

	// Setup Protected Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Login Successful")
	})

	app.Listen(":3000")
}

// EOF
