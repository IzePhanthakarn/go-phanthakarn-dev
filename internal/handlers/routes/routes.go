package routes

import (
	ctx "context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/utils"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/handlers/middlewares"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/guest"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/healthcheck"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/me"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/user"
	swagger "github.com/arsmn/fiber-swagger/v2"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/sirupsen/logrus"
)

const (
	// MaximumSize100MB body limit 100 mb.
	MaximumSize100MB = 1024 * 1024 * 100
	// MaximumSize1MB body limit 1 mb.
	MaximumSize1MB = 1024 * 1024 * 1
)

// NewRouter new router
func NewRouter() {
	app := fiber.New(
		fiber.Config{
			IdleTimeout:    5 * time.Second,
			BodyLimit:      MaximumSize100MB,
			ReadBufferSize: MaximumSize1MB,
		},
	)

	app.Use(
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
		requestid.New(),
		cors.New(),
		middlewares.Logger(),
		middlewares.WrapError(),
	)
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Use(middlewares.AcceptLanguage())
	if config.CF.Swagger.Enable {
		v1.Get("/swagger/*", swagger.HandlerDefault)
	}

	// Waiting to use...
	requiredAuth := jwtware.New(jwtware.Config{
		Claims:        &context.Claims{},
		SigningMethod: jwt.SigningMethodES256.Name,
		SigningKey:    utils.VerifyKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(config.RR.Internal.Unauthorized.HTTPStatusCode()).
				JSON(config.RR.Internal.Unauthorized.WithLocale(c))
		},
	})

	requiredSuperAdmin := middlewares.RequiredRoles(models.SuperAdminRole)
	healthCheckEndpoint := healthcheck.NewEndpoint()
	guestEndpoint := guest.NewEndpoint()
	userEndpoint := user.NewEndpoint()
	meEndpoint := me.NewEndpoint()

	healthCheck := v1.Group("health-check")
	{
		healthCheck.Get("/", healthCheckEndpoint.HealthCheck)
	}

	member := v1.Group("/c")
	{
		guest := member.Group("/guest")
		{
			guest.Post("/register", guestEndpoint.Register)
			guest.Post("/login", guestEndpoint.Login)
			guest.Post("/refresh-token", guestEndpoint.RefreshToken)
			guest.Post("/reset-password", guestEndpoint.ResetPassword)
		}

		me := member.Group("/me", requiredAuth)
		{
			me.Get("", meEndpoint.GetMe)
			me.Post("/logout", meEndpoint.Logout)
		}
	}

	// doctor := v1.Group("/d", requiredAuth, requiredDoctor)
	// {

	// }

	admin := v1.Group("/a", requiredAuth, requiredSuperAdmin)
	{
		users := admin.Group("/users", requiredAuth)
		{
			users.Get("", userEndpoint.GetAll)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
		defer cancel()

		logrus.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	logrus.Infof("Approved server on port: %d ...", config.CF.App.Port)
	err := app.Listen(fmt.Sprintf(":%d", config.CF.App.Port))
	if err != nil {
		logrus.Panic(err)
	}
}
