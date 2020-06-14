package server

import (
	"encoding/json"
	"instagram_bot/bot"
	"instagram_bot/database"
	"io/ioutil"

	"github.com/gofiber/fiber"
	"github.com/gofiber/session"
	"github.com/spf13/viper"
)

// Server holds the fiber instance
type Server struct {
	app     *fiber.App
	db      *database.Database
	session *session.Session
}

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// New will return a new Server instance and wraps all the routess
func New(db *database.Database) *Server {
	app := fiber.New()
	session := session.New()

	s := &Server{app, db, session}

	s.apiRoutes()
	s.authRoutes()
	s.frontRoutes()

	s.app.Post("/webhook", func(c *fiber.Ctx) {
		r := new(bot.WebhookResponse)
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
	s.app.Listen(viper.Get("server_port"))
}

func respond(c *fiber.Ctx, res response) {
	if err := c.JSON(res); err != nil {
		c.Status(500).Send(err)
	}
}
