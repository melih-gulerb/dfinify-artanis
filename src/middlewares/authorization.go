package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type JWTClaims struct {
	OrganizationID string `json:"organizationId"`
	jwt.RegisteredClaims
}

func AuthorizationMiddleware(secretKey string) fiber.Handler {
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

		tokenStr := parts[1]

		token, err := parseJWTToken(tokenStr, secretKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fmt.Sprintf("invalid token: %s", err.Error()),
			})
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		c.Locals("organizationId", claims.OrganizationID)

		return c.Next()
	}
}

func parseJWTToken(tokenStr string, secretKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

func ExtractOrganizationID(c *fiber.Ctx) (string, error) {
	organizationID := c.Locals("organizationId")
	if organizationID == nil {
		return "", errors.New("organizationId not found in context")
	}

	id, ok := organizationID.(string)
	if !ok {
		return "", errors.New("organizationId has invalid type")
	}

	return id, nil
}
