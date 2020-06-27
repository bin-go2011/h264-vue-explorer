package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Static("/", "../frontend/dist")

	app.Listen(3000)
}
