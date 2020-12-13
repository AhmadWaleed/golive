package main

import (
	"github.com/brendonferreira/golive"
	components "github.com/brendonferreira/golive/examples/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	liveServer := golive.NewServer()

	app.Get("/clock", liveServer.CreateHTMLHandler(components.NewClock, golive.PageContent{
		Lang:  "us",
		Title: "Hello world",
	}))
	app.Get("/todo", liveServer.CreateHTMLHandler(components.NewTodo, golive.PageContent{
		Lang:  "us",
		Title: "Hello world",
	}))

	app.Get("/slider", liveServer.CreateHTMLHandler(components.NewSlider, golive.PageContent{
		Lang:  "us",
		Title: "Hello world",
	}))

	app.Get("/ws", websocket.New(liveServer.HandleWSRequest))

	_ = app.Listen(":3000")

}
