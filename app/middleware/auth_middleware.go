package middleware

import (
	"simak-go/app/config"
	"simak-go/app/helper"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Protected(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Missing Authorization Header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid Token Format")
		}

		tokenString := parts[1]
		claims, err := helper.ParseToken(tokenString, cfg.JWTSecret)
		if err != nil {
			return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or Expired Token")
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role").(string)
		for _, role := range allowedRoles {
			if role == userRole {
				return c.Next()
			}
		}
		return helper.ErrorResponse(c, fiber.StatusForbidden, "Access Denied")
	}
}
