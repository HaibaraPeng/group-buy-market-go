package infrastructure

import (
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain"
	"time"
)

// MySQLProductRepository is a GORM implementation of ProductRepository
type MySQLProductRepository struct {
	db *gorm.DB
}

// NewMySQLProductRepository creates a new MySQL product repository
func NewMySQLProductRepository(db *gorm.DB) *MySQLProductRepository {
	return &MySQLProductRepository{
		db: db,
	}
}

// Save saves a product
func (r *MySQLProductRepository) Save(product *domain.Product) error {
	return r.db.Save(product).Error
}

// FindByID finds a product by ID
func (r *MySQLProductRepository) FindByID(id int64) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// FindAll returns all products
func (r *MySQLProductRepository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

// MySQLGroupBuyActivityRepository is a GORM implementation of GroupBuyActivityRepository
type MySQLGroupBuyActivityRepository struct {
	db *gorm.DB
}

// NewMySQLGroupBuyActivityRepository creates a new MySQL group buy activity repository
func NewMySQLGroupBuyActivityRepository(db *gorm.DB) *MySQLGroupBuyActivityRepository {
	return &MySQLGroupBuyActivityRepository{
		db: db,
	}
}

// Save saves a group buy activity
func (r *MySQLGroupBuyActivityRepository) Save(activity *domain.GroupBuyActivity) error {
	return r.db.Save(activity).Error
}

// FindByID finds a group buy activity by ID
func (r *MySQLGroupBuyActivityRepository) FindByID(id int64) (*domain.GroupBuyActivity, error) {
	var activity domain.GroupBuyActivity
	err := r.db.First(&activity, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}

// FindByActivityID finds a group buy activity by activity ID
func (r *MySQLGroupBuyActivityRepository) FindByActivityID(activityID int64) (*domain.GroupBuyActivity, error) {
	var activity domain.GroupBuyActivity
	err := r.db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}

// FindAll returns all group buy activities
func (r *MySQLGroupBuyActivityRepository) FindAll() ([]*domain.GroupBuyActivity, error) {
	var activities []*domain.GroupBuyActivity
	err := r.db.Find(&activities).Error
	return activities, err
}

// UpdateStatus updates the status of a group buy activity
func (r *MySQLGroupBuyActivityRepository) UpdateStatus(id int64, status int) error {
	return r.db.Model(&domain.GroupBuyActivity{}).Where("id = ?", id).Update("status", status).Update("update_time", time.Now()).Error
}
