package model

import "time"

// DiscountTypeEnum 折扣类型枚举
type DiscountTypeEnum byte

const (
	BaseDiscount DiscountTypeEnum = iota // 基础折扣
	TagDiscount                          // 标签折扣
)

// MarketPlanEnum 营销计划枚举
type MarketPlanEnum string

const (
	ZJ MarketPlanEnum = "ZJ" // 直减
	ZK MarketPlanEnum = "ZK" // 折扣
	MJ MarketPlanEnum = "MJ" // 满减
	N  MarketPlanEnum = "N"  // N元购
)

// GroupBuyDiscountVO represents the discount part of GroupBuyActivityDiscountVO
type GroupBuyDiscountVO struct {
	DiscountName string           `json:"discount_name"`
	DiscountDesc string           `json:"discount_desc"`
	DiscountType DiscountTypeEnum `json:"discount_type"`
	MarketPlan   MarketPlanEnum   `json:"market_plan"`
	MarketExpr   string           `json:"market_expr"`
	TagId        string           `json:"tag_id"`
}

// GroupBuyActivityDiscountVO represents the combined view of group buy activity and discount
type GroupBuyActivityDiscountVO struct {
	ActivityId       int64               `json:"activity_id"`
	ActivityName     string              `json:"activity_name"`
	Source           string              `json:"source"`
	Channel          string              `json:"channel"`
	GoodsId          string              `json:"goods_id"`
	GroupBuyDiscount *GroupBuyDiscountVO `json:"group_buy_discount"`
	GroupType        int                 `json:"group_type"`
	TakeLimitCount   int                 `json:"take_limit_count"`
	Target           int                 `json:"target"`
	ValidTime        int                 `json:"valid_time"`
	Status           int                 `json:"status"`
	StartTime        time.Time           `json:"start_time"`
	EndTime          time.Time           `json:"end_time"`
	TagId            string              `json:"tag_id"`
	TagScope         string              `json:"tag_scope"`
}

// SkuVO represents the SKU value object
type SkuVO struct {
	// Goods ID
	GoodsId string `json:"goods_id"`
	// Goods name
	GoodsName string `json:"goods_name"`
	// Original price
	OriginalPrice float64 `json:"original_price"`
}

// SCSkuActivityVO represents the SC SKU activity value object
type SCSkuActivityVO struct {
	// Channel
	Source string `json:"source"`
	// Channel
	Channel string `json:"channel"`
	// Activity ID
	ActivityId int64 `json:"activity_id"`
	// Goods ID
	GoodsId string `json:"goods_id"`
}
