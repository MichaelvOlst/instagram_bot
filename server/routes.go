package server

import (
	"instagram_bot/bot"

	"github.com/gofiber/fiber"
)

func (s *Server) setupAPIRoutes() {

	api := s.app.Group("/api")

	api.Post("/instagram/posts", func(c *fiber.Ctx) {

		r := new(bot.InstagramRequest)
		// Parse body into struct
		if err := c.BodyParser(r); err != nil {
			c.Status(400).Send(err)
			return
		}

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
