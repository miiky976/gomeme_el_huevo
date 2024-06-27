package handler

import (
	"bufio"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var omega = make(chan uint)

func Stream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	ctx := c.Context()

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		fmt.Println("WRITER")
		for {
			select {
			case msg := <-omega:
				fmt.Fprintf(w, "data: Message: %d\n\n", msg)
				fmt.Println(msg)
				err := w.Flush()
				if err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					return
				}
			case <-ctx.Done():
				fmt.Println("Client disconnected")
				return
			}
		}
	}))
	return nil
}
