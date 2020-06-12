package server

import (
	"net/http"

	"github.com/gofiber/fiber"
)

// authMiddleware
func (s *Server) authMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) {

		token := ctx.Get("api_token")
		_, err := s.db.CheckToken(token)
		if err != nil {
			ctx.JSON(response{Error: err.Error()})
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
