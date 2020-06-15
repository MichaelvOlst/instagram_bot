package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
)

// apiMiddleware
func (s *Server) apiMiddleware() fiber.Handler {
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

// authMiddleware
func (s *Server) authMiddleware() fiber.Handler {

	return func(ctx *fiber.Ctx) {

		store := s.session.Get(ctx)
		defer store.Save()

		authID := store.Get("auth_id")

		fmt.Println(authID)

		if authID == nil {
			ctx.Status(401)
			return
		}

		_, err := s.db.GetUserByID(authID.(int64))
		if err != nil {
			ctx.Status(401)
			return
		}

		ctx.Next()
	}
}
