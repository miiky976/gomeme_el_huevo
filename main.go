package main

import (
	"miiky976/Godis/handler"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	app.Static("/", "./public")

	app.Get("/GETLOOP/:place", handler.GetLast)

	app.Get("/GETALL", handler.GetAll)

	app.Post("/SET", handler.AddText)

	app.Post("/SETFILE", handler.AddFile)

	app.Listen(port)
}
