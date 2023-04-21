package guest

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/bcrypt"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/utils"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/token"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/pkg/user"
	"github.com/sirupsen/logrus"
)

// Service this is a guest service interface
type ServiceInterface interface {
	Register(c *context.Context, f *RegisterForm) (*models.Token, error)
	Login(c *context.Context, f *LoginForm) (*models.Token, error)
	ResetPassword(c *context.Context, f *ResetPasswordForm) error
}

type service struct {
	config    *config.Configs
	result    *config.ReturnResult
	userRepo  user.RepositoryInterface
	tokenServ token.ServiceInterface
}

// NewService new a guest service
func NewService(config *config.Configs, result *config.ReturnResult) ServiceInterface {
	return &service{
		config:    config,
		result:    result,
		userRepo:  user.NewRepository(),
		tokenServ: token.NewService(config, result),
	}
}

// Register register
func (s *service) Register(c *context.Context, f *RegisterForm) (*models.Token, error) {
	birthday, err := utils.StringToDate(f.Birthday)
	if err != nil {
		return nil, err
	}

	user, _ := s.userRepo.FindOneByCitizenID(c.GetDatabase(), f.CitizenID)
	if user != nil {
		return nil, s.result.Common.DuplicateCid
	}

	if f.Email != "" {
		user, _ := s.userRepo.FindOneByEmail(c.GetDatabase(), f.Email)
		if user != nil {
			return nil, s.result.Common.DuplicateEmail
		}
	}

	u := &models.User{
		Email:       f.Email,
		Prefix:      f.Prefix,
		FirstName:   f.FirstName,
		LastName:    f.LastName,
		PhoneNumber: f.PhoneNumber,
		CitizenID:   f.CitizenID,
		Birthday:    birthday,
	}

	ph, err := bcrypt.GeneratePassword(f.Password)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	u.PasswordHash = ph

	err = s.userRepo.Create(c.GetDatabase(), u)
	if err != nil {
		return nil, err
	}

	userRole := &models.UserRole{
		Role:   models.CustomerRole,
		UserID: u.ID,
	}
	err = s.userRepo.Create(c.GetDatabase(), userRole)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenServ.Create(c, u)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return token, nil
}

// Login login
func (s *service) Login(c *context.Context, f *LoginForm) (*models.Token, error) {
	u, _ := s.userRepo.FindOneByEmail(c.GetDatabase(), f.Username)
	if u == nil {
		u, _ = s.userRepo.FindOneByCitizenID(c.GetDatabase(), f.Username)
		if u == nil {
			return nil, s.result.Common.InvalidUsername
		}
	}

	if !bcrypt.ComparePassword(u.PasswordHash, f.Password) {
		return nil, s.result.Common.InvalidPassword
	}

	token, err := s.tokenServ.Create(c, u)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return token, nil
}
func (s *service) ResetPassword(c *context.Context, f *ResetPasswordForm) error {
	uq := user.Query{
		CitizenID:   f.CitizenID,
		Birthday:    f.Birthday,
		PhoneNumber: f.PhoneNumber,
	}
	us := &models.User{}
	if err := s.userRepo.FindOneByQuery(c.GetDatabase(), uq, us); err != nil {
		return s.result.Common.UserNotFound
	}

	ph, err := bcrypt.GeneratePassword(f.Password)
	if err != nil {
		logrus.Error(err)
		return err
	}
	us.PasswordHash = ph
	err = s.userRepo.Upsert(c.GetDatabase(), "id", []string{}, us)
	if err != nil {
		return err
	}
	return nil
}
