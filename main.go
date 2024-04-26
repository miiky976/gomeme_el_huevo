package main

import (
	"miiky976/Godis/handler"
	"miiky976/Godis/kv"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	app.Static("/", "./public")

	app.Get("/templ", func(c *fiber.Ctx) error {
		kv.Test()
		return c.SendString("kv template loaded :)")
	})

	app.Post("/SET", handler.AddText)

	app.Post("/SETFILE", handler.AddFile)

	app.Listen(port)
}
