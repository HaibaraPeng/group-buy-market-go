package po

import (
	"time"
)

// CrowdTagsJob represents a crowd tagging job entity
type CrowdTagsJob struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Tag ID
	TagId string `json:"tag_id" gorm:"column:tag_id"`
	// Batch ID
	BatchId string `json:"batch_id" gorm:"column:batch_id"`
	// Tag type (participation volume, consumption amount)
	TagType int `json:"tag_type" gorm:"column:tag_type"`
	// Tag rule (limit type N times)
	TagRule string `json:"tag_rule" gorm:"column:tag_rule"`
	// Statistics data, start time
	StatStartTime time.Time `json:"stat_start_time" gorm:"column:stat_start_time"`
	// Statistics data, end time
	StatEndTime time.Time `json:"stat_end_time" gorm:"column:stat_end_time"`
	// Status; 0 initial, 1 planned (entering execution phase), 2 reset, 3 completed
	Status int `json:"status" gorm:"column:status"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for CrowdTagsJob
func (CrowdTagsJob) TableName() string {
	return "crowd_tags_job"
}
