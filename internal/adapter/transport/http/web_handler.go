package http

import (
	"github.com/gofiber/fiber/v3"
	"time"
)

var year = time.Now().Year()

func (s *server) HomePageHandler(c fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title":    "Test",
		"Subtitle": "Test-1",
	})

}
