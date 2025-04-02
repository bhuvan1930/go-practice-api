package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello from Go Practice API!")
    })

    app.Get("/weather", func(c *fiber.Ctx) error {
        forecast := []string{"Sunny", "Cloudy", "Rainy"}
        return c.JSON(forecast)
    })

    app.Listen(":3000")
}
