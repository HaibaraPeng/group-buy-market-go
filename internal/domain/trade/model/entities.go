package model

import "time"

// MarketPayOrderEntity 拼团，预购订单营销实体对象
type MarketPayOrderEntity struct {
	// 预购订单ID
	OrderId string `json:"orderId"`
	// 折扣金额
	DeductionPrice float64 `json:"deductionPrice"`
	// 交易订单状态枚举
	TradeOrderStatusEnumVO TradeOrderStatusEnumVO `json:"tradeOrderStatusEnumVO"`
}

// UserEntity 用户实体
type UserEntity struct {
	UserId string `json:"userId"`
}

// PayActivityEntity 支付活动实体
type PayActivityEntity struct {
	TeamId       string    `json:"teamId"`
	ActivityId   int64     `json:"activityId"`
	ActivityName string    `json:"activityName"`
	TargetCount  int       `json:"targetCount"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}

// PayDiscountEntity 支付折扣实体
type PayDiscountEntity struct {
	Source         string  `json:"source"`
	Channel        string  `json:"channel"`
	GoodsId        string  `json:"goodsId"`
	GoodsName      string  `json:"goodsName"`
	OriginalPrice  float64 `json:"originalPrice"`
	DeductionPrice float64 `json:"deductionPrice"`
	OutTradeNo     string  `json:"outTradeNo"`
}
