package route

import (
	"simak-go/app/config"
	"simak-go/app/handler"
	"simak-go/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, authHandler *handler.AuthHandler) {
	// Global Middleware
	app.Use(logger.New())

	// API Group
	api := app.Group("/api")

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Example Protected Route
	user := api.Group("/users", middleware.Protected(cfg))
	user.Get("/profile", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "User Profile",
			"data": fiber.Map{
				"user_id": userID,
				"role":    role,
			},
		})
	})

	// Example Admin Only Route
	admin := api.Group("/admin", middleware.Protected(cfg), middleware.RoleMiddleware("admin"))
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "Welcome Admin",
		})
	})
}
