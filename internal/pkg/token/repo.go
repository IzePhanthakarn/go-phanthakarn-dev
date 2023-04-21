package token

import (
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/repositories"
	"gorm.io/gorm"
)

// RepositoryInterface repo interface
type RepositoryInterface interface {
	Create(db *gorm.DB, i interface{}) error
	Delete(db *gorm.DB, i interface{}) error
	Update(db *gorm.DB, i interface{}) error
	FindOneByID(db *gorm.DB, id uint) (*models.Token, error)
	FindOneByRefreshToken(db *gorm.DB, token string) (*models.Token, error)
}

// Repository repo
type Repository struct {
	repositories.Repository
}

// NewRepo new repo
func NewRepo() *Repository {
	return &Repository{
		repositories.NewRepository(),
	}
}

// FindOneByID find one
func (r *Repository) FindOneByID(db *gorm.DB, id uint) (*models.Token, error) {
	t := &models.Token{}
	db = db.Preload("User")
	if err := r.FindOneObjectByIDUInt(db, id, t); err != nil {
		return nil, err
	}
	return t, nil
}

// FindOneByRefreshToken find one by refresh token
func (r *Repository) FindOneByRefreshToken(db *gorm.DB, token string) (*models.Token, error) {
	t := &models.Token{}
	err := db.
		Preload("User").
		Preload("User.Roles").
		Where("refresh_token = ?", token).
		First(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
