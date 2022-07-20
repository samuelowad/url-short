package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/url-shorterner/database"
	"github.com/samuelowad/url-shorterner/router"
)

func main() {
	app := fiber.New()

	database.Init()
	router.SetupRouter(app)

	app.Listen(":3000")
}
