package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Promotion struct {
	Id             uint
	ExternalId     string
	Price          float32
	ExpirationDate time.Time
}

func main() {
	// db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=promotion port=5432"), &gorm.Config{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world !")
	})

	app.Listen(":3000")
}
