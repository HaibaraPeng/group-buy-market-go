package po

import (
	"time"
)

// CrowdTags represents a crowd tag entity
type CrowdTags struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Crowd tag ID
	TagId string `json:"tag_id" gorm:"column:tag_id"`
	// Crowd tag name
	TagName string `json:"tag_name" gorm:"column:tag_name"`
	// Crowd tag description
	TagDesc string `json:"tag_desc" gorm:"column:tag_desc"`
	// Crowd tag statistics
	Statistics int `json:"statistics" gorm:"column:statistics"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// TableName specifies the table name for CrowdTags
func (CrowdTags) TableName() string {
	return "crowd_tags"
}
