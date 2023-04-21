package repositories

import (
	"context"
	"fmt"
	"sync"

	"math"
	"time"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Repository common repository
type Repository struct {
	mutex sync.Mutex
}

// NewRepository new repository
func NewRepository() Repository {
	return Repository{}
}

// DefaultContext default context
func (r *Repository) DefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)
}

// FindOneObjectByID find one
func (r *Repository) FindOneObjectByID(db *gorm.DB, id uint64, i interface{}) error {
	return r.FindOneObjectByField(db, "id", id, i)
}

// FindOneObjectByIDInt find one
func (r *Repository) FindOneObjectByIDInt(db *gorm.DB, id int, i interface{}) error {
	return r.FindOneObjectByField(db, "id", id, i)

}

// FindOneObjectByIDUInt find one
func (r *Repository) FindOneObjectByIDUInt(db *gorm.DB, id uint, i interface{}) error {
	return r.FindOneObjectByField(db, "id", id, i)
}

// FindOneObjectByUserIDUInt find one
func (r *Repository) FindOneObjectByUserIDUInt(db *gorm.DB, id uint, i interface{}) error {
	return r.FindOneObjectByField(db, "user_id", id, i)
}

// FindOneObjectByIDString find one
func (r *Repository) FindOneObjectByIDString(db *gorm.DB, field string, value string, i interface{}) error {
	return r.FindOneObjectByField(db, field, value, i)
}

// FindOneObjectByField find one
func (r *Repository) FindOneObjectByField(db *gorm.DB, field string, value interface{}, i interface{}) error {
	return db.Where(fmt.Sprintf("%s = ?", field), value).First(i).Error
}

// FindOneLastObjectByField find one
func (r *Repository) FindOneLastObjectByField(db *gorm.DB, field string, value interface{}, i interface{}) error {
	return db.Where(fmt.Sprintf("%s = ?", field), value).Last(i).Error
}

// Create create
func (r *Repository) Create(db *gorm.DB, i interface{}) error {
	return db.Omit(clause.Associations).Create(i).Error
}

// CreateInBatch create with batch size
func (r *Repository) CreateInBatch(db *gorm.DB, i interface{}, batchSize int) error {
	return db.Omit(clause.Associations).CreateInBatches(i, batchSize).Error
}

// CreateWithAssociation create with association
func (r *Repository) CreateWithAssociation(db *gorm.DB, i interface{}) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Save(i).Error
}

// Update update
func (r *Repository) Update(db *gorm.DB, i interface{}) error {
	return db.Omit(clause.Associations).Save(i).Error
}

// Delete update stamp deleted_at
func (r *Repository) Delete(db *gorm.DB, i interface{}) error {
	return db.Omit(clause.Associations).Delete(i).Error
}

// DeleteWithCondition delete with condition
func (r *Repository) DeleteWithCondition(db *gorm.DB, field string, value, i interface{}) error {
	return db.Omit(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).Delete(i).Error
}

// FindOneByIDFullAssociations find one by id full associations
func (r *Repository) FindOneByIDFullAssociations(db *gorm.DB, id uint64, i interface{}) error {
	return r.FindOneObjectByID(db.Preload(clause.Associations), id, i)
}

// FindOneByFieldFullAssociations find one full associations
func (r *Repository) FindOneByFieldFullAssociations(db *gorm.DB, field string, value interface{}, i interface{}) error {
	return db.Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), value).First(i).Error
}

// FindAllByIDs get all by ids
func (r *Repository) FindAllByIDs(db *gorm.DB, ids []uint, i interface{}) error {
	return db.Where("id in (?)", ids).Find(i).Error
}

// FindAllByStrings get all by strins
func (r *Repository) FindAllByStrings(db *gorm.DB, field string, values []string, i interface{}) error {
	return db.Where(fmt.Sprintf("%s in (?)", field), values).Find(i).Error
}

// FindAllByField get all by field
func (r *Repository) FindAllByField(db *gorm.DB, field string, values interface{}, i interface{}) error {
	return db.Where(fmt.Sprintf("%s in (?)", field), values).Find(i).Error
}

// FindAllByFieldFullAssociations get all by field
func (r *Repository) FindAllByFieldFullAssociations(db *gorm.DB, field string, values interface{}, i interface{}) error {
	return db.Preload(clause.Associations).Where(fmt.Sprintf("%s = ?", field), values).Find(i).Error
}

// FindAllByValues get all by values
func (r *Repository) FindAllByValues(db *gorm.DB, field string, values interface{}, i interface{}) error {
	return db.Where(fmt.Sprintf("%s IN (?)", field), values).Find(i).Error
}

// FindByOldID find by old id
func (r *Repository) FindByOldID(db *gorm.DB, id interface{}, i interface{}) error {
	return db.Where("old_id = ?", id, i).First(i).Error
}

// Upsert upsert
func (r *Repository) Upsert(db *gorm.DB, uniqueKey string, columns []string, i interface{}) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: uniqueKey}},
		DoUpdates: clause.AssignmentColumns(columns),
		UpdateAll: len(columns) == 0,
	}).
		Omit(clause.Associations).
		Create(i).Error
}

// BulkUpsert bulk upsert
func (r *Repository) BulkUpsert(db *gorm.DB, uniqueKeys []string, updateColumns []string, i interface{}, batchSize int) error {
	uk := []clause.Column{}
	for _, uniqueKey := range uniqueKeys {
		uk = append(uk, clause.Column{Name: uniqueKey})
	}
	return db.Clauses(clause.OnConflict{
		Columns:   uk,
		DoUpdates: clause.AssignmentColumns(updateColumns),
		UpdateAll: len(updateColumns) == 0,
	}).
		Omit(clause.Associations).
		CreateInBatches(i, batchSize).Error
}

// SoftDelete soft delete
func (r *Repository) SoftDelete(db *gorm.DB, field string, value, actorID, i interface{}) error {
	values := map[string]interface{}{
		"deleted_by_employee_id": actorID,
		"deleted_at":             time.Now(),
	}
	err := db.
		Model(i).
		Where(fmt.Sprintf("%s = ?", field), value).
		Updates(values).Error
	if err != nil {
		return err
	}

	return nil
}

// PageForm page info interface
type PageForm interface {
	GetPage() int
	GetSize() int
	GetQuery() string
	GetSort() string
	GetReverse() bool
	GetOrderBy() string
}

type PagePatientForm interface {
	GetPage() int
	GetSize() int
	GetQuery() string
	GetHn() string
	GetFname() string
	GetCid() string
	GetStartDate() string
	GetEndDate() string
	GetPatientColorId() string
	GetSort() string
	GetReverse() bool
	GetOrderBy() string
}

const (
	// DefaultPage default page in page query
	DefaultPage int = 1
	// DefaultSize default size in page query
	DefaultSize int = 20
)

// FindAllAndPageInformation get page information
func (r *Repository) FindAllAndPageInformation(db *gorm.DB, pageForm PageForm, entities interface{}, selector ...string) (*models.PageInformation, error) {
	var count int64
	db.Model(entities).Count(&count)

	if pageForm.GetOrderBy() != "" {
		db = db.Order(pageForm.GetOrderBy())
	} else if pageForm.GetSort() != "" {
		order := pageForm.GetSort()
		if pageForm.GetReverse() {
			order = order + " desc"
		}
		db = db.Order(order)
	} else {
		db = db.Order("id")
	}

	if pageForm.GetQuery() != "" {
		db = db.Where(pageForm.GetQuery())
	}

	page := pageForm.GetPage()
	if pageForm.GetPage() < 1 {
		page = DefaultPage
	}

	limit := pageForm.GetSize()
	if pageForm.GetSize() == 0 {
		limit = DefaultSize
	}

	var offset int
	if page != 1 {
		offset = (page - 1) * limit
	}

	if len(selector) > 0 {
		db = db.Select(selector)
	}

	if err := db.
		Limit(limit).
		Offset(offset).
		Find(entities).Error; err != nil {
		return nil, err
	}

	return &models.PageInformation{
		Page:     page,
		Size:     limit,
		Count:    count,
		LastPage: int(math.Ceil(float64(count) / float64(limit))),
	}, nil
}
