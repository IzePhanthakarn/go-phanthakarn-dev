package middlewares

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"

	"github.com/gofiber/fiber/v2"
)

const (
	// EN english language
	EN = "en"

	// TH thai language
	TH = "th"
)

// AcceptLanguage header Accept-Language
func AcceptLanguage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := c.Get(fiber.HeaderAcceptLanguage)
		if lang != EN && lang != TH {
			lang = EN
		}

		c.Locals(context.LangKey, lang)
		return c.Next()
	}
}
