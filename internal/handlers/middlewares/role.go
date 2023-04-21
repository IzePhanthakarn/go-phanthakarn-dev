package middlewares

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/handlers/render"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/gofiber/fiber/v2"
)

// RequiredRoles required roles
func RequiredRoles(roles ...models.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.New(c)
		claim := ctx.GetClaims()
		for _, role := range roles {
			for _, userRole := range claim.Roles {
				if models.Role(userRole) == role {
					return c.Next()
				}
			}
		}
		return render.Error(c, fiber.ErrUnauthorized)
	}
}
