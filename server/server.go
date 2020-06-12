package server

import (
	"encoding/json"
	"instagram_bot/bot"
	"instagram_bot/database"
	"io/ioutil"

	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
)

// Server holds the fiber instance
type Server struct {
	app *fiber.App
	db  *database.Database
}

type response struct {
	Data  interface{} `json:",omitempty"`
	Error interface{} `json:",omitempty"`
}

// New will return a new Server instance and wraps all the routess
func New(db *database.Database) *Server {
	app := fiber.New()

	s := &Server{app, db}

	// basicAuthCfg := basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "123456",
	// 	},
	// }
	// s.app.Use(basicauth.New(basicAuthCfg))
	s.app.Use(s.authMiddleware())

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
	s.app.Listen(viper.Get("server_port"))
}
