package po

import (
	"time"
)

// GroupBuyOrder represents a group buying order entity
type GroupBuyOrder struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Team ID
	TeamId string `json:"team_id" gorm:"column:team_id"`
	// Activity ID
	ActivityId int64 `json:"activity_id" gorm:"column:activity_id"`
	// Source
	Source string `json:"source" gorm:"column:source"`
	// Channel
	Channel string `json:"channel" gorm:"column:channel"`
	// Original price
	OriginalPrice float64 `json:"original_price" gorm:"column:original_price"`
	// Deduction price
	DeductionPrice float64 `json:"deduction_price" gorm:"column:deduction_price"`
	// Payment price
	PayPrice float64 `json:"pay_price" gorm:"column:pay_price"`
	// Target count
	TargetCount int `json:"target_count" gorm:"column:target_count"`
	// Completed count
	CompleteCount int `json:"complete_count" gorm:"column:complete_count"`
	// Locked count
	LockCount int `json:"lock_count" gorm:"column:lock_count"`
	// Status (0-pending, 1-completed, 2-failed)
	Status int `json:"status" gorm:"column:status"`
	// Valid start time for group buying
	ValidStartTime time.Time `json:"valid_start_time" gorm:"column:valid_start_time"`
	// Valid end time for group buying
	ValidEndTime time.Time `json:"valid_end_time" gorm:"column:valid_end_time"`
	// Notify URL
	NotifyUrl string `json:"notify_url" gorm:"column:notify_url"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for GroupBuyOrder
func (GroupBuyOrder) TableName() string {
	return "group_buy_order"
}
