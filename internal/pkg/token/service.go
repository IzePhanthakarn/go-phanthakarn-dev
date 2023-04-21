package token

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/context"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/utils"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

// ServiceInterface service interface
type ServiceInterface interface {
	Create(c *context.Context, u *models.User) (*models.Token, error)
	RefreshToken(c *context.Context, f *RefeshTokenForm) (*models.Token, error)
}

type service struct {
	jwtKey    *ecdsa.PrivateKey
	tokenRepo RepositoryInterface
}

// NewService new service
func NewService(config *config.Configs, result *config.ReturnResult) *service {
	return &service{
		jwtKey:    utils.SignKey,
		tokenRepo: NewRepo(),
	}
}

// Create create
func (s *service) Create(c *context.Context, u *models.User) (*models.Token, error) {
	t, err := s.createJWTToken(c, u)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return t, nil
}

func (s *service) createJWTToken(c *context.Context, u *models.User) (*models.Token, error) {
	t := &models.Token{
		UserID: u.ID,
		User:   u,
	}
	rto, err := s.createRefreshToken(u)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	t.RefreshToken = rto

	claims := &context.Claims{}
	now := time.Now()
	claims.Subject = fmt.Sprint(u.ID)
	claims.Issuer = "pwa"
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Add(24 * time.Hour).Unix()
	err = s.tokenRepo.Create(c.GetDatabase(), t)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	claims.RefreshTokenID = t.ID
	claims.Roles = u.RoleValues()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return nil, err
	}
	t.Token = tokenString
	return t, nil
}

func (s *service) createRefreshToken(u *models.User) (string, error) {
	rts := fmt.Sprintf("%d%s", u.ID, time.Now().String())
	h := sha1.New()
	_, err := h.Write([]byte(rts))
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// RefreshToken refresh token
func (s *service) RefreshToken(c *context.Context, f *RefeshTokenForm) (*models.Token, error) {
	t, err := s.tokenRepo.FindOneByRefreshToken(c.GetDatabase(), f.RefreshToken)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = s.tokenRepo.Delete(c.GetDatabase(), t)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	t, err = s.Create(c, t.User)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return t, nil
}
