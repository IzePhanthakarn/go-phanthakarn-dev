package guest

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/handlers"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
}

type endpoint struct {
	config       *config.Configs
	result       *config.ReturnResult
	serv         ServiceInterface
	tokenService token.ServiceInterface
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config:       config.CF,
		result:       config.RR,
		serv:         NewService(config.CF, config.RR),
		tokenService: token.NewService(config.CF, config.RR),
	}
}

// Register register
// @Tags Authentication
// @Summary Register
// @Description Register
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body RegisterForm true "body for register"
// @Success 200 {object} models.Token
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /c/guest/register [post]
func (ep *endpoint) Register(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.serv.Register, &RegisterForm{})
}

// Login login
// @Tags Authentication
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body LoginForm true "body for login"
// @Success 200 {object} models.Token
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /c/guest/login [post]
func (ep *endpoint) Login(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.serv.Login, &LoginForm{})
}

// RefreshToken refresh token
// @Tags Authentication
// @Summary refresh token
// @Description refresh token
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body token.RefeshTokenForm true "body for refresh token"
// @Success 200 {object} models.Token
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /c/guest/refresh-token [post]
func (ep *endpoint) RefreshToken(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.tokenService.RefreshToken, &token.RefeshTokenForm{})
}

// ResetPassword reset password
// @Tags Authentication
// @Summary Reset password
// @Description Reset password
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body ResetPasswordForm true "body for reset password"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /c/guest/reset-password [post]
func (ep *endpoint) ResetPassword(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, ep.serv.ResetPassword, &ResetPasswordForm{})
}
