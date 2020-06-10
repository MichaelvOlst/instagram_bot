package server

import (
	"fmt"

	"github.com/gofiber/fiber"
)

// authMiddleware
func (s *Server) authMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				ctx.Next(err)
				return
			}
		}()
		ctx.Next()
	}
}
