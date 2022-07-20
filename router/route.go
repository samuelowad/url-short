package router

import (
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/url-shorterner/database"
	"github.com/samuelowad/url-shorterner/models"
)

func SetupRouter(app *fiber.App) {

	app.Post("/v1/create-link", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		var data models.Model

		if err := c.BodyParser(&data); err != nil {
			return err
		}
		_, err := net.LookupIP(data.Url)
		// if cant resolve ip, return error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Could not find inputted domain",
			})
		}

		data.Short = randString()

		err = database.CreateUrl(&data)

		data.Short = "http://localhost:3000/v1/" + data.Short

		if err != nil {
			log.Fatal(err)
		}

		// err = database.CreateUrl(&data)

		return c.Status(fiber.StatusCreated).JSON(data)
	})

	app.Get("/v1/:short", func(c *fiber.Ctx) error {
		short := c.Params("short")

		data, err := database.FindUrl(short)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Could not find inputted short url",
			})
		}
		return c.RedirectBack("http://" + data.Url)
	})

}

func randString() string {

	var shorturl string
	rand.Seed(time.Now().Unix())

	str := "abcdefghijklmnopqrstuvwxyz1234567890-"

	shuff := []rune(str)

	// Shuffling the string
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})
	shorturl = string(shuff)

	data, _ := database.FindUrl(shorturl)

	if data.Short != "" {

		return randString()
	}

	return shorturl

}
