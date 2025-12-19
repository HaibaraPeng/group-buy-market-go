package po

import (
	"time"
)

// CrowdTagsDetail represents a crowd tag detail entity
type CrowdTagsDetail struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Crowd tag ID
	TagId string `json:"tag_id" gorm:"column:tag_id"`
	// User ID
	UserId string `json:"user_id" gorm:"column:user_id"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for CrowdTagsDetail
func (CrowdTagsDetail) TableName() string {
	return "crowd_tags_detail"
}
