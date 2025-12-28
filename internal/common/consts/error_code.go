package consts

// ResponseCode 定义错误代码常量
type ResponseCode string

const (
	// E0103 用户参与次数超限错误
	E0103 ResponseCode = "E0103"
)

// GetErrorMessage 根据错误代码获取错误信息
func (rc ResponseCode) GetErrorMessage() string {
	switch rc {
	case E0103:
		return "用户参与次数已达上限"
	default:
		return "未知错误"
	}
}
