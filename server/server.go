package server

import (
	"encoding/json"
	"instagram_bot/bot"
	"instagram_bot/config"
	"io/ioutil"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

// Server holds the fiber instance
type Server struct {
	app *fiber.App
	cfg *config.Config
}

// New will return a new Server instance and wraps all the routess
func New(cfg *config.Config) *Server {
	app := fiber.New()

	s := &Server{app, cfg}

	basicAuthCfg := basicauth.Config{
		Users: map[string]string{
			"admin": "123456",
		},
	}
	s.app.Use(basicauth.New(basicAuthCfg))

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
	s.app.Listen(s.cfg.Server.Port)
}
