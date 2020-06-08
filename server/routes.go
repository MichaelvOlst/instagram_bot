package server

import (
	"fmt"
	"instagram_bot/bot"
	"log"

	"github.com/gofiber/fiber"
)

// Request struct
type Request struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Profile  string `json: "profile"`
	Webhook  string `json: "webhook"`
}

func (s *Server) setupAPIRoutes() {

	api := s.app.Group("/api")

	api.Post("/instagram/posts", func(c *fiber.Ctx) {

		r := new(Request)
		// Parse body into struct
		if err := c.BodyParser(r); err != nil {
			c.Status(400).Send(err)
			return
		}

		b := bot.New(r.Username, r.Password, r.Profile, r.Webhook)

		go func() {

			urls, err := b.GetPosts()
			if err != nil {
				log.Fatal(err)
			}

			for _, result := range urls {
				fmt.Println(result)
				// fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
			}
		}()

		// Return Employees in JSON format
		if err := c.JSON(r); err != nil {
			c.Status(500).Send(err)
			return
		}
	})
}
