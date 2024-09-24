package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterPostRoutes(apiRouter fiber.Router) {
	postRouter := apiRouter.Group("/posts")

	postRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("post route")
	})
}
