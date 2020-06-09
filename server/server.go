package server

import (
	"encoding/json"
	"instagram_bot/bot"
	"io/ioutil"

	"github.com/gofiber/fiber"
)

// Server holds the fiber instance
type Server struct {
	app *fiber.App
}

// New will return a new Server instance and wraps all the routess
func New() *Server {
	app := fiber.New()

	s := &Server{app}

	s.setupAPIRoutes()

	s.app.Post("/webhook", func(c *fiber.Ctx) {
		r := new(bot.WebhookResponse)
		// Parse body into struct
		if err := c.BodyParser(r); err != nil {
			c.Status(400).Send(err)
			return
		}

		bytes, _ := json.Marshal(r)
		err := ioutil.WriteFile("test.json", bytes, 0777)

		if err != nil {
			c.Status(400).Send(err)
			return
		}
	})

	return s
}

// Start will listen to the api requests
func (s *Server) Start() {
	s.app.Listen(3000)
}
