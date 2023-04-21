package render

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"

	"github.com/gofiber/fiber/v2"
)

// JSON render json to client
func JSON(c *fiber.Ctx, response interface{}) error {
	return c.
		Status(config.RR.Internal.Success.HTTPStatusCode()).
		JSON(response)
}

// Byte render byte to client
func Byte(c *fiber.Ctx, bytes []byte) error {
	_, err := c.Status(config.RR.Internal.Success.HTTPStatusCode()).
		Write(bytes)

	return err
}

// Error render error to client
func Error(c *fiber.Ctx, err error) error {
	if locErr, ok := err.(config.Result); ok {
		return c.
			Status(locErr.HTTPStatusCode()).
			JSON(locErr.WithLocale(c))
	}

	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.
			Status(fiberErr.Code).
			JSON(config.NewResultWithMessage(fiberErr.Error()))
	}

	defaultErr := config.RR.Internal.General
	return c.
		Status(defaultErr.HTTPStatusCode()).
		JSON(defaultErr.WithLocale(c))
}
