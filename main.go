package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sugarcodex/simple-api/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.HomeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen(port)
}
