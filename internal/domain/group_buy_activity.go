package domain

import (
	"time"
)

// GroupBuyActivity represents a group buying activity entity
type GroupBuyActivity struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Activity ID
	ActivityId int64 `json:"activity_id" gorm:"column:activity_id"`
	// Activity name
	ActivityName string `json:"activity_name" gorm:"column:activity_name"`
	// Source
	Source string `json:"source" gorm:"column:source"`
	// Channel
	Channel string `json:"channel" gorm:"column:channel"`
	// Goods ID
	GoodsId string `json:"goods_id" gorm:"column:goods_id"`
	// Discount ID
	DiscountId string `json:"discount_id" gorm:"column:discount_id"`
	// Group type (0 auto-group, 1 target-group)
	GroupType int `json:"group_type" gorm:"column:group_type"`
	// Take limit count
	TakeLimitCount int `json:"take_limit_count" gorm:"column:take_limit_count"`
	// Target
	Target int `json:"target" gorm:"column:target"`
	// Valid time (minutes)
	ValidTime int `json:"valid_time" gorm:"column:valid_time"`
	// Status (0 created, 1 active, 2 expired, 3 discarded)
	Status int `json:"status" gorm:"column:status"`
	// Activity start time
	StartTime time.Time `json:"start_time" gorm:"column:start_time"`
	// Activity end time
	EndTime time.Time `json:"end_time" gorm:"column:end_time"`
	// Tag ID
	TagId string `json:"tag_id" gorm:"column:tag_id"`
	// Tag scope
	TagScope string `json:"tag_scope" gorm:"column:tag_scope"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// TableName specifies the table name for GroupBuyActivity
func (GroupBuyActivity) TableName() string {
	return "group_buy_activities"
}
