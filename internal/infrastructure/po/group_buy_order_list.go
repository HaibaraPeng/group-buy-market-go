package po

import (
	"time"
)

// GroupBuyOrderList represents a group buying order list entity
type GroupBuyOrderList struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// User ID
	UserId string `json:"user_id" gorm:"column:user_id"`
	// Team ID
	TeamId string `json:"team_id" gorm:"column:team_id"`
	// Order ID
	OrderId string `json:"order_id" gorm:"column:order_id"`
	// Activity ID
	ActivityId int64 `json:"activity_id" gorm:"column:activity_id"`
	// Activity start time
	StartTime time.Time `json:"start_time" gorm:"column:start_time"`
	// Activity end time
	EndTime time.Time `json:"end_time" gorm:"column:end_time"`
	// Goods ID
	GoodsId string `json:"goods_id" gorm:"column:goods_id"`
	// Source
	Source string `json:"source" gorm:"column:source"`
	// Channel
	Channel string `json:"channel" gorm:"column:channel"`
	// Original price
	OriginalPrice float64 `json:"original_price" gorm:"column:original_price"`
	// Deduction price
	DeductionPrice float64 `json:"deduction_price" gorm:"column:deduction_price"`
	// Pay price
	PayPrice float64 `json:"pay_price" gorm:"column:pay_price"`
	// Status (0-initial locked, 1-consumption completed)
	Status int `json:"status" gorm:"column:status"`
	// External trade number - ensures external call uniqueness
	OutTradeNo string `json:"out_trade_no" gorm:"column:out_trade_no"`
	// External trade time
	OutTradeTime *time.Time `json:"out_trade_time" gorm:"column:out_trade_time"`
	// 唯一业务ID
	BizId string `json:"biz_id" gorm:"column:biz_id"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
	// Count for limit operations
	Count int64 `json:"count" gorm:"-"`
}

// TableName specifies the table name for GroupBuyOrderList
func (GroupBuyOrderList) TableName() string {
	return "group_buy_order_list"
}
