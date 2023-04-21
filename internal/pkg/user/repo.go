package user

import (
	"fmt"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/repositories"
	"gorm.io/gorm"
)

// Repository repo interface
type RepositoryInterface interface {
	Create(db *gorm.DB, i interface{}) error
	FindAll(db *gorm.DB, f *GetAllForm) (*models.Page, error)
	FindOneByID(db *gorm.DB, id uint, i interface{}) error
	FindOneByPatientID(db *gorm.DB, id uint) (*models.User, error)
	FindOneByCitizenID(db *gorm.DB, cid string) (*models.User, error)
	FindOneByEmail(db *gorm.DB, email string) (*models.User, error)
	FindOneByQuery(db *gorm.DB, query Query, i interface{}) error
	FindOneByQuery2(db *gorm.DB, query Query2, i interface{}) error
	Upsert(db *gorm.DB, uniqueKey string, columns []string, i interface{}) error
	FindOneObjectByIDUInt(db *gorm.DB, id uint, i interface{}) error
}

type repository struct {
	repositories.Repository
}

// NewRepository new repository
func NewRepository() RepositoryInterface {
	return &repository{
		repositories.NewRepository(),
	}
}

type Query struct {
	CitizenID   string
	Birthday    string
	PhoneNumber string
}

type Query2 struct {
	Id int
}

func preload(db *gorm.DB) *gorm.DB {
	return db.Preload("Roles")
}

// FindOneByID find one by id
func (r *repository) FindOneByID(db *gorm.DB, id uint, i interface{}) error {
	db = preload(db)
	return r.FindOneObjectByIDUInt(db, id, i)
}

// FindOneByPatientID find one by patient id
func (r *repository) FindOneByPatientID(db *gorm.DB, id uint) (*models.User, error) {
	user := &models.User{}
	err := r.FindOneObjectByField(db, "patient_id", id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindAll find all users
func (r *repository) FindAll(db *gorm.DB, f *GetAllForm) (*models.Page, error) {
	entities := []*models.User{}
	if f.Role != nil {
		db = db.Joins("INNER JOIN user_roles on user_roles.user_id = users.id").
			Where("user_roles.role = ?", f.Role)
	}
	p, err := r.FindAllAndPageInformation(db, &f.PageForm, &entities)
	if err != nil {
		return nil, err
	}
	return models.NewPage(p, entities), nil
}

// FindOneByCitizenID find user by citizen id
func (r *repository) FindOneByEmail(db *gorm.DB, email string) (*models.User, error) {
	fmt.Println("eiei")
	user := &models.User{}
	fmt.Println("eiei2")
	// db = preload(db)
	fmt.Println("eiei3")
	err := r.FindOneObjectByField(db, "email", email, user)
	fmt.Println("eiei4")
	if err != nil {
		fmt.Println("eiei5")
		return nil, err
	}
	fmt.Println("eiei6")
	return user, nil
}

// FindOneByCitizenID find user by citizen id

func (r *repository) FindOneByCitizenID(db *gorm.DB, cid string) (*models.User, error) {
	user := &models.User{}
	db = preload(db)
	err := r.FindOneObjectByField(db, "cid", cid, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindOneByQuery find one by query
func (r *repository) FindOneByQuery(db *gorm.DB, query Query, i interface{}) error {
	db = db.Where("cid = ? AND birthday = ? AND phone_number = ?",
		query.CitizenID,
		query.Birthday,
		query.PhoneNumber)

	return db.First(i).Error
}

// FindOneByQuery2 find one by query2
func (r *repository) FindOneByQuery2(db *gorm.DB, query Query2, i interface{}) error {
	db = db.Where("id = ?",
		query.Id)
	return db.First(i).Error
}
