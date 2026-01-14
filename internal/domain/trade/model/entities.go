package model

import "time"

// MarketPayOrderEntity 拼团，预购订单营销实体对象
type MarketPayOrderEntity struct {
	// 团队ID
	TeamId string `json:"teamId"`
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
	ValidTime    int       `json:"validTime"`
}

// PayDiscountEntity 支付折扣实体
type PayDiscountEntity struct {
	Source         string  `json:"source"`
	Channel        string  `json:"channel"`
	GoodsId        string  `json:"goodsId"`
	GoodsName      string  `json:"goodsName"`
	OriginalPrice  float64 `json:"originalPrice"`
	DeductionPrice float64 `json:"deductionPrice"`
	PayPrice       float64 `json:"payPrice"`
	OutTradeNo     string  `json:"outTradeNo"`
	NotifyUrl      string  `json:"notifyUrl"`
}

// GroupBuyActivityEntity 拼团活动实体对象
type GroupBuyActivityEntity struct {
	// 活动ID
	ActivityId int64 `json:"activityId"`
	// 活动名称
	ActivityName string `json:"activityName"`
	// 折扣ID
	DiscountId string `json:"discountId"`
	// 拼团方式（0自动成团、1达成目标拼团）
	GroupType int `json:"groupType"`
	// 拼团次数限制
	TakeLimitCount int `json:"takeLimitCount"`
	// 拼团目标
	Target int `json:"target"`
	// 拼团时长（分钟）
	ValidTime int `json:"validTime"`
	// 活动状态（0创建、1生效、2过期、3废弃）
	Status ActivityStatusEnumVO `json:"status"`
	// 活动开始时间
	StartTime time.Time `json:"startTime"`
	// 活动结束时间
	EndTime time.Time `json:"endTime"`
	// 人群标签规则标识
	TagId string `json:"tagId"`
	// 人群标签规则范围
	TagScope string `json:"tagScope"`
}

// TradeRuleCommandEntity 拼团交易命令实体
type TradeRuleCommandEntity struct {
	// 用户ID
	UserId string `json:"userId"`
	// 活动ID
	ActivityId int64 `json:"activityId"`
}

// TradeRuleFilterBackEntity 拼团交易，过滤反馈实体
type TradeRuleFilterBackEntity struct {
	// 用户参与活动的订单量
	UserTakeOrderCount int `json:"userTakeOrderCount"`
}

// GroupBuyTeamEntity 拼团团队实体
type GroupBuyTeamEntity struct {
	// 团队ID
	TeamId string `json:"teamId"`
	// 活动ID
	ActivityId int64 `json:"activityId"`
	// 目标数量
	TargetCount int `json:"targetCount"`
	// 完成数量
	CompleteCount int `json:"completeCount"`
	// 锁定数量
	LockCount int `json:"lockCount"`
	// 状态
	Status GroupBuyOrderEnumVO `json:"status"`
	// 拼团开始时间 - 参与拼团时间
	ValidStartTime time.Time `json:"validStartTime"`
	// 拼团结束时间 - 拼团有效时长
	ValidEndTime time.Time `json:"validEndTime"`
	// 回调配置
	NotifyConfigVO *NotifyConfigVO `json:"notifyConfigVO"`
}

// TradePaySuccessEntity 交易支付成功实体
type TradePaySuccessEntity struct {
	// 渠道
	Source string `json:"source"`
	// 来源
	Channel string `json:"channel"`
	// 用户ID
	UserId string `json:"userId"`
	// 外部交易号
	OutTradeNo string `json:"outTradeNo"`
	// 外部交易时间
	OutTradeTime time.Time `json:"outTradeTime"`
}

// GroupBuyTeamSettlementAggregate 拼团团队结算聚合根
type GroupBuyTeamSettlementAggregate struct {
	UserEntity            *UserEntity            `json:"userEntity"`
	GroupBuyTeamEntity    *GroupBuyTeamEntity    `json:"groupBuyTeamEntity"`
	TradePaySuccessEntity *TradePaySuccessEntity `json:"tradePaySuccessEntity"`
}

// TradePaySettlementEntity 交易支付结算实体
type TradePaySettlementEntity struct {
	Source     string `json:"source"`
	Channel    string `json:"channel"`
	UserId     string `json:"userId"`
	TeamId     string `json:"teamId"`
	ActivityId int64  `json:"activityId"`
	OutTradeNo string `json:"outTradeNo"`
}

// TradeSettlementRuleCommandEntity 拼团交易结算规则命令
type TradeSettlementRuleCommandEntity struct {
	// 渠道
	Source string `json:"source"`
	// 来源
	Channel string `json:"channel"`
	// 用户ID
	UserId string `json:"userId"`
	// 外部交易单号
	OutTradeNo string `json:"outTradeNo"`
	// 外部交易时间
	OutTradeTime time.Time `json:"outTradeTime"`
}

// TradeSettlementRuleFilterBackEntity 拼团交易结算规则反馈
type TradeSettlementRuleFilterBackEntity struct {
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
	// 状态（0-拼单中、1-完成、2-失败）
	Status GroupBuyOrderEnumVO `json:"status"`
	// 拼团开始时间 - 参与拼团时间
	ValidStartTime time.Time `json:"validStartTime"`
	// 拼团结束时间 - 拼团有效时长
	ValidEndTime time.Time `json:"validEndTime"`
	// 回调地址
	NotifyUrl string `json:"notifyUrl"`
}

// NotifyTaskEntity 回调任务实体
type NotifyTaskEntity struct {
	// 拼单组队ID
	TeamId string `json:"teamId"`
	// 回调接口
	NotifyUrl string `json:"notifyUrl"`
	// 回调次数
	NotifyCount int `json:"notifyCount"`
	// 参数对象
	ParameterJson string `json:"parameterJson"`
}

// LockKey 生成锁键
func (n *NotifyTaskEntity) LockKey() string {
	return "notify_job_lock_key_" + n.TeamId
}
