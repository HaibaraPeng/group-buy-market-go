package model

// MarketPayOrderEntity 拼团，预购订单营销实体对象
type MarketPayOrderEntity struct {
	// 预购订单ID
	OrderId string `json:"orderId"`
	// 折扣金额
	DeductionPrice float64 `json:"deductionPrice"`
	// 交易订单状态枚举
	TradeOrderStatusEnumVO TradeOrderStatusEnumVO `json:"tradeOrderStatusEnumVO"`
}
