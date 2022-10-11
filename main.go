package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sugarcodex/simple-api/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.HomeHandler)

	app.Listen(":3000")
}
