package server

import (
	"net/http"

	"github.com/gofiber/fiber"
)

// authMiddleware
func (s *Server) authMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) {

		_, err := s.db.CheckToken(ctx.Get("api_token"))
		if err != nil {
			ctx.JSON(response{Error: err.Error()})
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
