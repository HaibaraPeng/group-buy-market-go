package po

import (
	"time"
)

// Sku represents a SKU entity
type Sku struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Source
	Source string `json:"source" gorm:"column:source"`
	// Channel
	Channel string `json:"channel" gorm:"column:channel"`
	// Goods ID
	GoodsId string `json:"goods_id" gorm:"column:goods_id"`
	// Goods name
	GoodsName string `json:"goods_name" gorm:"column:goods_name"`
	// Original price
	OriginalPrice float64 `json:"original_price" gorm:"column:original_price"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for Sku
func (Sku) TableName() string {
	return "sku"
}
