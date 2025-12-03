package infrastructure

import (
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain"
	"time"
)

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
