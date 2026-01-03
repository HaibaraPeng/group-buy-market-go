package consts

// ResponseCode 定义错误代码常量
type ResponseCode string

const (
	// SUCCESS 成功
	SUCCESS ResponseCode = "0000"
	// UN_ERROR 未知失败
	UN_ERROR ResponseCode = "0001"
	// ILLEGAL_PARAMETER 非法参数
	ILLEGAL_PARAMETER ResponseCode = "0002"
	// INDEX_EXCEPTION 唯一索引冲突
	INDEX_EXCEPTION ResponseCode = "0003"
	// UPDATE_ZERO 更新记录为0
	UPDATE_ZERO ResponseCode = "0004"

	// E0001 不存在对应的折扣计算服务
	E0001 ResponseCode = "E0001"
	// E0002 无拼团营销配置
	E0002 ResponseCode = "E0002"
	// E0003 拼团活动降级拦截
	E0003 ResponseCode = "E0003"
	// E0004 拼团活动切量拦截
	E0004 ResponseCode = "E0004"
	// E0005 拼团组队失败，记录更新为0
	E0005 ResponseCode = "E0005"
	// E0006 拼团组队完结，锁单量已达成
	E0006 ResponseCode = "E0006"
	// E0007 拼团人群限定，不可参与
	E0007 ResponseCode = "E0007"

	// E0101 活动非生效状态错误
	E0101 ResponseCode = "E0101"
	// E0102 活动不在可参与时间范围内错误
	E0102 ResponseCode = "E0102"
	// E0103 用户参与次数超限错误
	E0103 ResponseCode = "E0103"
	// E0104 不存在的外部交易单号或用户已退单
	E0104 ResponseCode = "E0104"
	// E0105 SC渠道黑名单拦截
	E0105 ResponseCode = "E0105"
	// E0106 订单交易时间不在拼团有效时间范围内
	E0106 ResponseCode = "E0106"
)

// GetErrorMessage 根据错误代码获取错误信息
func (rc ResponseCode) GetErrorMessage() string {
	switch rc {
	case SUCCESS:
		return "成功"
	case UN_ERROR:
		return "未知失败"
	case ILLEGAL_PARAMETER:
		return "非法参数"
	case INDEX_EXCEPTION:
		return "唯一索引冲突"
	case UPDATE_ZERO:
		return "更新记录为0"
	case E0001:
		return "不存在对应的折扣计算服务"
	case E0002:
		return "无拼团营销配置"
	case E0003:
		return "拼团活动降级拦截"
	case E0004:
		return "拼团活动切量拦截"
	case E0005:
		return "拼团组队失败，记录更新为0"
	case E0006:
		return "拼团组队完结，锁单量已达成"
	case E0007:
		return "拼团人群限定，不可参与"
	case E0101:
		return "活动非生效状态"
	case E0102:
		return "活动不在可参与时间范围内"
	case E0103:
		return "用户参与次数已达上限"
	case E0104:
		return "不存在的外部交易单号或用户已退单"
	case E0105:
		return "SC渠道黑名单拦截"
	case E0106:
		return "订单交易时间不在拼团有效时间范围内"
	default:
		return "未知错误"
	}
}
