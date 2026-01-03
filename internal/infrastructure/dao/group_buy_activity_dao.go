package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
)

// GroupBuyActivityDAO defines the interface for group buy activity persistence
type GroupBuyActivityDAO interface {
	QueryGroupBuyActivityList(ctx context.Context) ([]*po.GroupBuyActivity, error)
	FindValidBySourceAndChannel(ctx context.Context, source string, channel string) (*po.GroupBuyActivity, error)
	FindValidByActivityID(ctx context.Context, activityID int64) (*po.GroupBuyActivity, error)
	QueryGroupBuyActivityByActivityId(ctx context.Context, activityID int64) (*po.GroupBuyActivity, error)
}

// MySQLGroupBuyActivityDAO is a GORM implementation of GroupBuyActivityDAO
type MySQLGroupBuyActivityDAO struct {
	data *data.Data
}

// NewMySQLGroupBuyActivityDAO creates a new MySQL group buy activity DAO
func NewMySQLGroupBuyActivityDAO(data *data.Data) GroupBuyActivityDAO {
	return &MySQLGroupBuyActivityDAO{
		data: data,
	}
}

// QueryGroupBuyActivityList returns all group buy activities
func (r *MySQLGroupBuyActivityDAO) QueryGroupBuyActivityList(ctx context.Context) ([]*po.GroupBuyActivity, error) {
	var activities []*po.GroupBuyActivity
	err := r.data.DB(ctx).WithContext(ctx).Find(&activities).Error
	return activities, err
}

// FindValidByActivityID finds a group buy activity by activity ID
func (r *MySQLGroupBuyActivityDAO) FindValidByActivityID(ctx context.Context, activityID int64) (*po.GroupBuyActivity, error) {
	var activity po.GroupBuyActivity
	err := r.data.DB(ctx).WithContext(ctx).Where("activity_id = ? AND status = ?", activityID, 1).First(&activity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}

// FindValidBySourceAndChannel finds the latest valid group buy activity by source and channel
func (r *MySQLGroupBuyActivityDAO) FindValidBySourceAndChannel(ctx context.Context, source string, channel string) (*po.GroupBuyActivity, error) {
	var activity po.GroupBuyActivity
	err := r.data.DB(ctx).WithContext(ctx).Where("source = ? AND channel = ?", source, channel).
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

// QueryGroupBuyActivityByActivityId finds a group buy activity by activity ID without considering status
func (r *MySQLGroupBuyActivityDAO) QueryGroupBuyActivityByActivityId(ctx context.Context, activityID int64) (*po.GroupBuyActivity, error) {
	var activity po.GroupBuyActivity
	err := r.data.DB(ctx).WithContext(ctx).Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}
