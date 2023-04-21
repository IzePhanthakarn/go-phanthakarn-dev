package me

import (
	"fmt"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"
	"github.com/sirupsen/logrus"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/client"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/token"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/user"
)

// ServiceInterface this is a user service interface
type ServiceInterface interface {
	GetMe(c *context.Context) (*models.User, error)
	Logout(c *context.Context) error
}

type service struct {
	config        *config.Configs
	result        *config.ReturnResult
	clientService client.Service
	userRepo      user.RepositoryInterface
	tokenRepo     token.RepositoryInterface
}

// NewService new a user service
func NewService(config *config.Configs, result *config.ReturnResult) ServiceInterface {
	return &service{
		config:        config,
		result:        result,
		clientService: client.NewService(config, result),
		userRepo:      user.NewRepository(),
		tokenRepo:     token.NewRepo(),
	}
}

// GetMe get me
func (s *service) GetMe(c *context.Context) (*models.User, error) {
	uid := c.GetUserID()
	fmt.Println(uid)
	u := &models.User{}
	if err := s.userRepo.FindOneByID(c.GetDatabase(), uid, u); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return u, nil
}

// Logout logout
func (s *service) Logout(c *context.Context) error {
	rtid := c.GetClaims().RefreshTokenID

	token, err := s.tokenRepo.FindOneByID(c.GetDatabase(), rtid)
	if err != nil {
		logrus.Error(err)
	}
	if err := s.tokenRepo.Delete(c.GetDatabase(), token); err != nil {
		logrus.Error(err)
	}

	return nil
}
