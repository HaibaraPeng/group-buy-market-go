package model

// MarketProductEntity 营销商品实体
// 表示参与营销活动的商品信息
type MarketProductEntity struct {
	// 活动ID
	ActivityId int64 `json:"activityId"`
	// 用户ID
	UserId string `json:"userId"`
	// 商品ID
	GoodsId string `json:"goodsId"`
	// 渠道
	Source string `json:"source"`
	// 来源
	Channel string `json:"channel"`
}

// TrialBalanceEntity 试算结果实体对象（给用户展示拼团可获得的优惠信息）
type TrialBalanceEntity struct {
	// 商品ID
	GoodsId string `json:"goodsId"`
	// 商品名称
	GoodsName string `json:"goodsName"`
	// 原始价格
	OriginalPrice float64 `json:"originalPrice"`
	// 折扣价格
	DeductionPrice float64 `json:"deductionPrice"`
	// 支付价格
	PayPrice float64 `json:"payPrice"`
	// 拼团目标数量
	TargetCount int `json:"targetCount"`
	// 拼团开始时间
	StartTime int64 `json:"startTime"`
	// 拼团结束时间
	EndTime int64 `json:"endTime"`
	// 是否可见拼团
	IsVisible bool `json:"isVisible"`
	// 是否可参与进团
	IsEnable bool `json:"isEnable"`
	// 拼团活动营销配置值对象
	GroupBuyActivityDiscountVO GroupBuyActivityDiscountVO `json:"groupBuyActivityDiscountVO"`
}
