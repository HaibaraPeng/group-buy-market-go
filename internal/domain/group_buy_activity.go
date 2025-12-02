package domain

import (
	"time"
)

// GroupBuyActivity represents a group buying activity entity
type GroupBuyActivity struct {
	// Auto-incrementing ID
	ID int64 `json:"id"`
	// Activity ID
	ActivityId int64 `json:"activity_id"`
	// Activity name
	ActivityName string `json:"activity_name"`
	// Source
	Source string `json:"source"`
	// Channel
	Channel string `json:"channel"`
	// Goods ID
	GoodsId string `json:"goods_id"`
	// Discount ID
	DiscountId string `json:"discount_id"`
	// Group type (0 auto-group, 1 target-group)
	GroupType int `json:"group_type"`
	// Take limit count
	TakeLimitCount int `json:"take_limit_count"`
	// Target
	Target int `json:"target"`
	// Valid time (minutes)
	ValidTime int `json:"valid_time"`
	// Status (0 created, 1 active, 2 expired, 3 discarded)
	Status int `json:"status"`
	// Activity start time
	StartTime time.Time `json:"start_time"`
	// Activity end time
	EndTime time.Time `json:"end_time"`
	// Tag ID
	TagId string `json:"tag_id"`
	// Tag scope
	TagScope string `json:"tag_scope"`
	// Creation time
	CreateTime time.Time `json:"create_time"`
	// Update time
	UpdateTime time.Time `json:"update_time"`
}
