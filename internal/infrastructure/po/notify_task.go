package po

import (
	"time"
)

// NotifyTask represents a notification callback task entity
type NotifyTask struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Activity ID
	ActivityId int64 `json:"activity_id" gorm:"column:activity_id"`
	// Group team ID
	TeamId string `json:"team_id" gorm:"column:team_id"`
	// Notify type (HTTP, MQ)
	NotifyType string `json:"notify_type" gorm:"column:notify_type"`
	// Notify message
	NotifyMQ string `json:"notify_mq" gorm:"column:notify_mq"`
	// Callback URL
	NotifyUrl string `json:"notify_url" gorm:"column:notify_url"`
	// Callback count
	NotifyCount int `json:"notify_count" gorm:"column:notify_count"`
	// Callback status [0-initial, 1-completed, 2-retry, 3-failed]
	NotifyStatus int `json:"notify_status" gorm:"column:notify_status"`
	// Parameter JSON object
	ParameterJson string `json:"parameter_json" gorm:"column:parameter_json"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for NotifyTask
func (NotifyTask) TableName() string {
	return "notify_task"
}
