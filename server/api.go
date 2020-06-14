package server

import (
	"encoding/base64"
	"instagram_bot/bot"

	"github.com/gofiber/fiber"
)

func (s *Server) apiRoutes() {

	api := s.app.Group("/api")
	api.Use(s.apiMiddleware())

	api.Post("/instagram/posts", func(c *fiber.Ctx) {

		r := new(bot.InstagramRequest)
		if err := c.BodyParser(r); err != nil {
			c.Status(400).Send(err)
			return
		}

		username, err := base64.StdEncoding.DecodeString(r.Username)
		if err != nil {
			c.Status(400).Send(err)
			return
		}

		password, err := base64.StdEncoding.DecodeString(r.Password)
		if err != nil {
			c.Status(400).Send(err)
			return
		}

		r.Username = string(username)
		r.Password = string(password)

		go func() {
			b := bot.New(r)
			b.Run()
		}()

		// Return Employees in JSON format
		if err := c.JSON(r); err != nil {
			c.Status(500).Send(err)
			return
		}
	})
}
