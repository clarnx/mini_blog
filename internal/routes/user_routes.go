package routes

import (
	"mini_blog/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/users")

	userRouter.Get("/", controllers.GetUsers)
}
