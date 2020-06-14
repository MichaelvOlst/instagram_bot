package server

import (
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

// // authMiddleware
// func (s *Server) authMiddleware() fiber.Handler {
// 	return func(ctx *fiber.Ctx) {

// 		store := s.session.Get(ctx) // get/create new session
// 		defer store.Save()

// 		authID := store.Get("auth_id")

// 		if authID == nil {
// 			ctx.Redirect("/login")
// 			return
// 		}

// 		u, err := s.db.GetUserByID(authID.(int64))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(u)

// 		// token := ctx.Get("api_token")
// 		// _, err := s.db.CheckToken(token)
// 		// if err != nil {
// 		// 	ctx.JSON(response{Error: err.Error()})
// 		// 	ctx.Status(http.StatusUnauthorized)
// 		// 	return
// 		// }

// 		ctx.Next()
// 	}
// }
