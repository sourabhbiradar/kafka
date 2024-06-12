package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()          // new fiber application
	api := app.Group("/api/v1") // grouping

	api.Post("/comments", createComment) // api/v1/comments

	app.Listen(":3000")
}
