package middlewares

import (
	"artanis/src/clients"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthorizationMiddleware(client *clients.DivineShieldClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing Authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid Authorization header format",
			})
		}

		user, err := client.Authorize(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "invalid or expired token",
				"details": err.Error(),
			})
		}

		c.Context().SetUserValue("user", user)

		return c.Next()
	}
}
