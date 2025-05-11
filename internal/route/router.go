package route

import (
	auth "go-fiber-app/internal/modules/auth"
	user "go-fiber-app/internal/modules/user"
	"go-fiber-app/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/auth/login", auth.LoginHandler)

	// Protected user routes
	userGroup := api.Group("/users", middleware.JWTMiddleware)
	userGroup.Post("/", user.CreateUserHandler)
	userGroup.Get("/", user.GetAllUsersHandler)
}
