package user

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/client"
)

// ServiceInterface this is a user service interface
type ServiceInterface interface {
	GetAll(c *context.Context, f *GetAllForm) (*models.Page, error)
}

type service struct {
	config        *config.Configs
	result        *config.ReturnResult
	clientService client.Service
	repo          RepositoryInterface
}

// NewService new a user service
func NewService(config *config.Configs, result *config.ReturnResult) ServiceInterface {
	return &service{
		config:        config,
		result:        result,
		clientService: client.NewService(config, result),
		repo:          NewRepository(),
	}
}

func (s *service) GetAll(c *context.Context, f *GetAllForm) (*models.Page, error) {
	return s.repo.FindAll(c.GetDatabase(), f)
}
