package server

import "github.com/gofiber/fiber"

// Server holds the fiber instance
type Server struct {
	app *fiber.App
}

// New will return a new Server instance and wraps all the routess
func New() *Server {
	app := fiber.New()

	s := &Server{app}

	s.setupAPIRoutes()

	return s
}

// Start will listen to the api requests
func (s *Server) Start() {
	s.app.Listen(3000)
}
