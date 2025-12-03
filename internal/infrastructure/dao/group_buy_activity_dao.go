package dao

import (
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// GroupBuyActivityDAO defines the interface for group buy activity persistence
type GroupBuyActivityDAO interface {
	Save(activity *po.GroupBuyActivity) error
	FindByID(id int64) (*po.GroupBuyActivity, error)
	FindByActivityID(activityID int64) (*po.GroupBuyActivity, error)
	FindAll() ([]*po.GroupBuyActivity, error)
	UpdateStatus(id int64, status int) error
}

// MySQLGroupBuyActivityDAO is a GORM implementation of GroupBuyActivityDAO
type MySQLGroupBuyActivityDAO struct {
	db *gorm.DB
}

// NewMySQLGroupBuyActivityDAO creates a new MySQL group buy activity DAO
func NewMySQLGroupBuyActivityDAO(db *gorm.DB) *MySQLGroupBuyActivityDAO {
	return &MySQLGroupBuyActivityDAO{
		db: db,
	}
}

// Save saves a group buy activity
func (r *MySQLGroupBuyActivityDAO) Save(activity *po.GroupBuyActivity) error {
	return r.db.Save(activity).Error
}

// FindByID finds a group buy activity by ID
func (r *MySQLGroupBuyActivityDAO) FindByID(id int64) (*po.GroupBuyActivity, error) {
	var activity po.GroupBuyActivity
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
func (r *MySQLGroupBuyActivityDAO) FindByActivityID(activityID int64) (*po.GroupBuyActivity, error) {
	var activity po.GroupBuyActivity
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
func (r *MySQLGroupBuyActivityDAO) FindAll() ([]*po.GroupBuyActivity, error) {
	var activities []*po.GroupBuyActivity
	err := r.db.Find(&activities).Error
	return activities, err
}

// UpdateStatus updates the status of a group buy activity
func (r *MySQLGroupBuyActivityDAO) UpdateStatus(id int64, status int) error {
	return r.db.Model(&po.GroupBuyActivity{}).Where("id = ?", id).Update("status", status).Update("update_time", time.Now()).Error
}
