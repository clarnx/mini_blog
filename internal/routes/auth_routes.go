package routes

import (
	"mini_blog/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(apiRouter fiber.Router) {
	authRouter := apiRouter.Group("/auth")

	authRouter.Post("/signup", controllers.SignUpUser)
}
