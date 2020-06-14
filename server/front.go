package server

import (
	"github.com/gofiber/fiber"
)

func (s *Server) frontRoutes() {

	s.app.Static("/", "public")

	auth := s.app.Group("")
	auth.Get("/*", func(ctx *fiber.Ctx) {
		ctx.SendFile("public/index.html")
	})
}
