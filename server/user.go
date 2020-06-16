package server

import (
	"instagram_bot/models"

	"github.com/gofiber/fiber"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) userRoutes() {

	g := s.app.Group("/api/users")
	g.Use(s.authMiddleware())
	g.Get("", s.getUsers)
	g.Post("", s.createUser)

}

func (s *Server) getUsers(c *fiber.Ctx) {

	result, err := s.db.GetUsers()
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	respond(c, response{Data: result})
}

func (s *Server) createUser(c *fiber.Ctx) {

	u := new(user)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}

	newUser := models.NewUser(u.Email, u.Password)

	err := s.db.CreateUser(&newUser)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	result, err := s.db.GetUsers()
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	respond(c, response{Data: result})
}
