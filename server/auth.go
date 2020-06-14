package server

import (
	"strings"

	"github.com/gofiber/fiber"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *login) Sanitize() {
	l.Email = strings.ToLower(strings.TrimSpace(l.Email))
}

func (s *Server) authRoutes() {

	auth := s.app.Group("/auth")
	auth.Post("/check", s.checkLogin)
	auth.Post("/login", s.login)
	auth.Post("/logout", s.logout)

}

func (s *Server) checkLogin(c *fiber.Ctx) {
	var err error
	store := s.session.Get(c)
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
}

func (s *Server) login(c *fiber.Ctx) {

	l := new(login)
	if err := c.BodyParser(l); err != nil {
		c.Status(400).Send(err)
		return
	}
	l.Sanitize()

	u, err := s.db.GetUserByEmail(l.Email)
	if err != nil {
		c.Status(422).Send(err)
		return
	}

	if err := u.ComparePassword(l.Password); err != nil {
		c.Status(422).Send(err)
		return
	}

	store := s.session.Get(c)
	defer store.Save()
	store.Set("auth_id", u.ID)

	respond(c, response{Data: u})
}

func (s *Server) logout(c *fiber.Ctx) {

	store := s.session.Get(c)
	defer store.Save()

	store.Delete("auth_id")

	respond(c, response{Data: true})
}
