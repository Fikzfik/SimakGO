package main

import (
	"simak-go/app/config"
	"simak-go/app/database"
	"simak-go/app/handler"
	"simak-go/app/repository"
	"simak-go/app/route"
	"simak-go/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect Database
	db := database.ConnectDB(cfg)

	// 3. Setup Dependencies (Dependency Injection)
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)

	// 4. Setup Fiber App
	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// 5. Setup Routes
	route.SetupRoutes(app, cfg, authHandler)

	// 6. Start Server
	err := app.Listen(":" + cfg.AppPort)
	if err != nil {
		panic(err)
	}
}
