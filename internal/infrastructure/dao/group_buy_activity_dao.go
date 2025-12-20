package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
)

// GroupBuyActivityDAO defines the interface for group buy activity persistence
type GroupBuyActivityDAO interface {
	QueryGroupBuyActivityList(ctx context.Context) ([]*po.GroupBuyActivity, error)
	FindValidBySourceAndChannel(source string, channel string) (*po.GroupBuyActivity, error)
}

// MySQLGroupBuyActivityDAO is a GORM implementation of GroupBuyActivityDAO
type MySQLGroupBuyActivityDAO struct {
	db *gorm.DB
}

// NewMySQLGroupBuyActivityDAO creates a new MySQL group buy activity DAO
func NewMySQLGroupBuyActivityDAO(db *gorm.DB) GroupBuyActivityDAO {
	return &MySQLGroupBuyActivityDAO{
		db: db,
	}
}

// QueryGroupBuyDiscountList returns all group buy activities
func (r *MySQLGroupBuyActivityDAO) QueryGroupBuyActivityList(ctx context.Context) ([]*po.GroupBuyActivity, error) {
	var activities []*po.GroupBuyActivity
	err := r.db.Find(&activities).Error
	return activities, err
}

// FindValidBySourceAndChannel finds the latest valid group buy activity by source and channel
func (r *MySQLGroupBuyActivityDAO) FindValidBySourceAndChannel(source string, channel string) (*po.GroupBuyActivity, error) {
	var activity po.GroupBuyActivity
	err := r.db.Where("source = ? AND channel = ?", source, channel).
		Order("id DESC").
		Limit(1).
		First(&activity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}
