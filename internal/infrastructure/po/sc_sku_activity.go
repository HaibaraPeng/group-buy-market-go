package po

import (
	"time"
)

// SCSkuActivity represents a channel product activity configuration association entity
type SCSkuActivity struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Channel
	Source string `json:"source" gorm:"column:source"`
	// Source
	Channel string `json:"channel" gorm:"column:channel"`
	// Activity ID
	ActivityId int64 `json:"activity_id" gorm:"column:activity_id"`
	// Goods ID
	GoodsId string `json:"goods_id" gorm:"column:goods_id"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for SCSkuActivity
func (SCSkuActivity) TableName() string {
	return "sc_sku_activity"
}
