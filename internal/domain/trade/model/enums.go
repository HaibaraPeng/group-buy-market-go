package model

// TradeOrderStatusEnumVO 交易订单状态枚举
type TradeOrderStatusEnumVO int

const (
	// CREATE 初始创建
	CREATE TradeOrderStatusEnumVO = iota

	// COMPLETE 消费完成
	COMPLETE

	// CLOSE 超时关单
	CLOSE
)

// String 返回状态枚举的字符串表示
func (t TradeOrderStatusEnumVO) String() string {
	switch t {
	case CREATE:
		return "初始创建"
	case COMPLETE:
		return "消费完成"
	case CLOSE:
		return "超时关单"
	default:
		return "未知状态"
	}
}

// Code 返回状态枚举的整数值
func (t TradeOrderStatusEnumVO) Code() int {
	return int(t)
}

// TradeOrderStatusEnumVOValueOf 根据整数值返回对应的状态枚举
func TradeOrderStatusEnumVOValueOf(code int) TradeOrderStatusEnumVO {
	switch code {
	case 0:
		return CREATE
	case 1:
		return COMPLETE
	case 2:
		return CLOSE
	default:
		return CREATE
	}
}

// ActivityStatusEnumVO 活动状态枚举
type ActivityStatusEnumVO int

const (
	// ACTIVITY_CREATE 创建
	ACTIVITY_CREATE ActivityStatusEnumVO = iota

	// ACTIVITY_ACTIVE 生效
	ACTIVITY_ACTIVE

	// ACTIVITY_EXPIRE 过期
	ACTIVITY_EXPIRE

	// ACTIVITY_DISCARD 废弃
	ACTIVITY_DISCARD
)

// String 返回活动状态枚举的字符串表示
func (a ActivityStatusEnumVO) String() string {
	switch a {
	case ACTIVITY_CREATE:
		return "创建"
	case ACTIVITY_ACTIVE:
		return "生效"
	case ACTIVITY_EXPIRE:
		return "过期"
	case ACTIVITY_DISCARD:
		return "废弃"
	default:
		return "未知状态"
	}
}

// Code 返回活动状态枚举的整数值
func (a ActivityStatusEnumVO) Code() int {
	return int(a)
}

// ActivityStatusEnumVOValueOf 根据整数值返回对应的活动状态枚举
func ActivityStatusEnumVOValueOf(code int) ActivityStatusEnumVO {
	switch code {
	case 0:
		return ACTIVITY_CREATE
	case 1:
		return ACTIVITY_ACTIVE
	case 2:
		return ACTIVITY_EXPIRE
	case 3:
		return ACTIVITY_DISCARD
	default:
		return ACTIVITY_CREATE
	}
}

// GroupBuyOrderEnumVO 拼团订单状态枚举
type GroupBuyOrderEnumVO int

const (
	// GROUP_BUY_PENDING 拼团中
	GROUP_BUY_PENDING GroupBuyOrderEnumVO = iota

	// GROUP_BUY_COMPLETE 拼团完成
	GROUP_BUY_COMPLETE

	// GROUP_BUY_FAILED 拼团失败
	GROUP_BUY_FAILED
)

// String 返回拼团订单状态枚举的字符串表示
func (g GroupBuyOrderEnumVO) String() string {
	switch g {
	case GROUP_BUY_PENDING:
		return "拼团中"
	case GROUP_BUY_COMPLETE:
		return "拼团完成"
	case GROUP_BUY_FAILED:
		return "拼团失败"
	default:
		return "未知状态"
	}
}

// Code 返回拼团订单状态枚举的整数值
func (g GroupBuyOrderEnumVO) Code() int {
	return int(g)
}

// GroupBuyOrderEnumVOValueOf 根据整数值返回对应的拼团订单状态枚举
func GroupBuyOrderEnumVOValueOf(code int) GroupBuyOrderEnumVO {
	switch code {
	case 0:
		return GROUP_BUY_PENDING
	case 1:
		return GROUP_BUY_COMPLETE
	case 2:
		return GROUP_BUY_FAILED
	default:
		return GROUP_BUY_PENDING
	}
}
