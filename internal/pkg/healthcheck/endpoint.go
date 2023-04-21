package healthcheck

import (
	"github.com/IzePhanthakarn/go-boilerplate/internal/core/config"
	"github.com/IzePhanthakarn/go-boilerplate/internal/handlers/render"
	"github.com/IzePhanthakarn/go-boilerplate/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	HealthCheck(c *fiber.Ctx) error
}

type endpoint struct {
	config *config.Configs
	result *config.ReturnResult
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config: config.CF,
		result: config.RR,
	}
}

// HealthCheck health check
// @Tags Payment
// @Summary health check
// @Description health check
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /health-check [get]
// @BasePath /api
func (ep *endpoint) HealthCheck(c *fiber.Ctx) error {
	return render.JSON(c, models.NewSuccessMessage())
}
