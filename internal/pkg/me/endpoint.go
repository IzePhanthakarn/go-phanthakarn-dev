package me

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	GetMe(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
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

// GetMe get me
// @Tags Me
// @Summary get me
// @Description get me
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} models.User
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /c/me [get]
func (ep *endpoint) GetMe(c *fiber.Ctx) error {
	return handlers.ResponseObjectWithoutRequest(c, ep.serv.GetMe)
}

// Logout logout
// @Tags Me
// @Summary logout
// @Description logout
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /c/me/logout [post]
func (ep *endpoint) Logout(c *fiber.Ctx) error {
	return handlers.ResponseSuccessWithoutRequest(c, ep.serv.Logout)
}
