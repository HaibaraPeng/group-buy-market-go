package po

import (
	"time"
)

// GroupBuyDiscount represents a group buying discount entity
type GroupBuyDiscount struct {
	// Auto-incrementing ID
	ID int64 `json:"id" gorm:"primaryKey;column:id"`
	// Discount ID
	DiscountId int `json:"discount_id" gorm:"column:discount_id"`
	// Discount title
	DiscountName string `json:"discount_name" gorm:"column:discount_name"`
	// Discount description
	DiscountDesc string `json:"discount_desc" gorm:"column:discount_desc"`
	// Discount type (0: base, 1: tag)
	DiscountType byte `json:"discount_type" gorm:"column:discount_type"`
	// Marketing plan (ZJ: direct reduction, MJ: full reduction, N yuan purchase)
	MarketPlan string `json:"market_plan" gorm:"column:market_plan"`
	// Marketing expression
	MarketExpr string `json:"market_expr" gorm:"column:market_expr"`
	// Crowd tag, specific offer limitation
	TagId string `json:"tag_id" gorm:"column:tag_id"`
	// Creation time
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	// Update time
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

// TableName specifies the table name for GroupBuyDiscount
func (GroupBuyDiscount) TableName() string {
	return "group_buy_discount"
}
