package handler

import (
	"crypto/rand"
	"fmt"
	"miiky976/Godis/kv"
	"miiky976/Godis/templates"
	"os"
	"strconv"
	"time"

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
	noteb := []byte(note)
	kv.Create(noteb, "string")
	return Render(c, templates.Noti("success", "Text saved"))
}

func AddFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return Render(c, templates.Noti("failed", "Fail on upload"))
	}
	if err := c.SaveFile(file, "/tmp/"+file.Filename); err != nil {
		return Render(c, templates.Noti("failed", "Fail on save"))
	}
	head := file.Header.Get("Content-Type")
	osfile, _ := os.ReadFile("/tmp/" + file.Filename)
	os.Remove("/tmp/" + file.Filename)
	kv.Create(osfile, head)
	return Render(c, templates.Noti("success", "File saved"))
}

func GetLast(c *fiber.Ctx) error {
	place, _ := strconv.Atoi(c.Params("place", "0"))
	if place >= len(kv.Master) {
		return Render(c, templates.Loop(place))
	}
	return Render(c, templates.Joiner(place))
}

// experimental
func Stream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	for true {
		msg := generateKey()
		c.SendString("Hola: " + msg)

		time.Sleep(1 * time.Second)
	}
	return nil
}
