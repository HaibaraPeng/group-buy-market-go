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
