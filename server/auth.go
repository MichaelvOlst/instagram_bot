package server

import (
	"github.com/gofiber/fiber"
)

func (s *Server) authRoutes() {

	auth := s.app.Group("/auth")
	auth.Post("/check", func(c *fiber.Ctx) {

		var err error
		store := s.session.Get(c) // get/create new session
		defer store.Save()

		authID := store.Get("auth_id")
		if authID == nil {
			respond(c, response{Data: false})
			return
		}

		u, err := s.db.GetUserByID(authID.(int64))
		if err != nil {
			respond(c, response{Error: err})
			return
		}

		respond(c, response{Data: u})
	})
}
