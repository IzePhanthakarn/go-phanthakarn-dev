package middlewares

import (
	"github.com/IzePhanthakarn/go-boilerplate/internal/core/context"
	"github.com/gofiber/fiber/v2"
)

const (
	authHeader = "Authorization"
)

// RequireAuthentication require authentication
func RequireAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get session token from authorization header
		token := c.Get(authHeader)

		// Add the user session to locals
		c.Locals(context.SessionTokenKey, token)
		return c.Next()
	}
}
