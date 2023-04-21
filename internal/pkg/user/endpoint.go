package user

import (
	"github.com/IzePhanthakarn/go-boilerplate/internal/core/config"
	"github.com/IzePhanthakarn/go-boilerplate/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	GetAll(c *fiber.Ctx) error
}

type endpoint struct {
	config *config.Configs
	result *config.ReturnResult
	serv   ServiceInterface
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config: config.CF,
		result: config.RR,
		serv:   NewService(config.CF, config.RR),
	}
}

// GetAll get all
// @Tags Users
// @Summary GetAll
// @Description GetAll
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request query GetAllForm true "query for get all users"
// @Success 200 {object} models.Page
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /a/users [get]
func (ep *endpoint) GetAll(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.serv.GetAll, &GetAllForm{})
}
