package handler

import (
	"crypto/rand"
	"fmt"
	"log"
	"miiky976/Godis/kv"
	"miiky976/Godis/templates"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(c *fiber.Ctx, com templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(com))(c)
}

func generateKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func AddText(c *fiber.Ctx) error {
	note := c.FormValue("note")
	key := generateKey()
	kv.SET(key, note, "string")
	return Render(c, templates.Text(key, note))
}

func AddFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	key := generateKey()
	contentType := string(c.Request().Header.ContentType())
	log.Printf("Content-Type: %s", contentType)
	if err != nil {
		log.Println("No se subio", err)
		return c.SendString("No se subio :( <br>")
	}
	c.SaveFile(file, "/tmp/"+file.Filename)
	head := file.Header.Get("Content-Type")
	osfile, _ := os.ReadFile("/tmp/" + file.Filename)
	kv.SET(key, osfile, head)
	return Render(c, templates.Image(key, head, osfile))
}

func GetAll(c *fiber.Ctx) error {
	return Render(c, templates.All())
}
