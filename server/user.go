package server

import (
	"github.com/gofiber/fiber"
)

func (s *Server) userRoutes() {

	g := s.app.Group("/api/users")
	g.Use(s.authMiddleware())
	g.Get("", s.getUsers)

}

func (s *Server) getUsers(c *fiber.Ctx) {
	// var err error
	// store := s.session.Get(c)
	// defer store.Save()

	// authID := store.Get("auth_id")
	// if authID == nil {
	// 	respond(c, response{Data: false})
	// 	return
	// }

	// u, err := s.db.GetUserByID(authID.(int64))
	// if err != nil {
	// 	respond(c, response{Error: err})
	// 	return

	// }

	respond(c, response{Data: true})
}
