package model

import (
	"group-buy-market-go/internal/common/consts"
	"strings"
	"time"
)

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

// TagScopeEnumVO 标签范围枚举
type TagScopeEnumVO struct {
	allow  bool
	refuse bool
	desc   string
}

var (
	// VISIBLE 可见性枚举 - 是否可看见拼团
	VISIBLE = &TagScopeEnumVO{allow: true, refuse: false, desc: "是否可看见拼团"}

	// ENABLE 可用性枚举 - 是否可参与拼团
	ENABLE = &TagScopeEnumVO{allow: true, refuse: false, desc: "是否可参与拼团"}
)

// Allow 获取允许状态
func (t *TagScopeEnumVO) Allow() bool {
	return t.allow
}

// Refuse 获取拒绝状态
func (t *TagScopeEnumVO) Refuse() bool {
	return t.refuse
}

// Desc 获取描述
func (t *TagScopeEnumVO) Desc() string {
	return t.desc
}

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

// IsVisible 可见限制
// 只要存在这样一个值，那么首次获得的默认值就是 false
func (g *GroupBuyActivityDiscountVO) IsVisible() bool {
	if g.TagScope == "" {
		return VISIBLE.Allow()
	}

	split := strings.Split(g.TagScope, consts.SPLIT)
	if len(split) > 0 && split[0] == "1" {
		return VISIBLE.Refuse()
	}

	return VISIBLE.Allow()
}

// IsEnable 参与限制
// 只要存在这样一个值，那么首次获得的默认值就是 false
func (g *GroupBuyActivityDiscountVO) IsEnable() bool {
	if g.TagScope == "" {
		return ENABLE.Allow()
	}

	split := strings.Split(g.TagScope, consts.SPLIT)
	if len(split) == 2 && split[1] == "2" {
		return ENABLE.Refuse()
	}

	return ENABLE.Allow()
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
