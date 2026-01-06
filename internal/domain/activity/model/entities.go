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

// UserGroupBuyOrderDetailEntity 拼团组队实体对象
type UserGroupBuyOrderDetailEntity struct {
	// 用户ID
	UserId string `json:"userId"`
	// 拼单组队ID
	TeamId string `json:"teamId"`
	// 活动ID
	ActivityId int64 `json:"activityId"`
	// 目标数量
	TargetCount int `json:"targetCount"`
	// 完成数量
	CompleteCount int `json:"completeCount"`
	// 锁单数量
	LockCount int `json:"lockCount"`
	// 拼团开始时间 - 参与拼团时间
	ValidStartTime int64 `json:"validStartTime"`
	// 拼团结束时间 - 拼团有效时长
	ValidEndTime int64 `json:"validEndTime"`
	// 外部交易单号-确保外部调用唯一幂等
	OutTradeNo string `json:"outTradeNo"`
}
