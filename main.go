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

	kv.Test()

	app.Static("/", "./public")

	app.Get("/GETLOOP/:place", handler.GetLast)

	app.Post("/SET", handler.AddText)

	app.Post("/SETFILE", handler.AddFile)

	app.Get("/load/:place", handler.Load)

	app.Get("/loadall", handler.LoadAll)

	app.Listen(port)
}
