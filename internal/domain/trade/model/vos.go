package model

// GroupBuyProgressVO 拼团进度值对象
type GroupBuyProgressVO struct {
	CompleteCount int `json:"completeCount"`
	TargetCount   int `json:"targetCount"`
	LockCount     int `json:"lockCount"`
}

// NotifyTaskHTTPEnumVO 拼团回调任务HTTP枚举值对象
// 对应Java中的NotifyTaskHTTPEnumVO枚举
type NotifyTaskHTTPEnumVO string

const (
	// SUCCESS 成功
	SUCCESS NotifyTaskHTTPEnumVO = "success"
	// ERROR 错误
	ERROR NotifyTaskHTTPEnumVO = "error"
	// NULL 空值
	NULL NotifyTaskHTTPEnumVO = "1001"
)

// NotifyTypeEnumVO 回调类型枚举值对象
type NotifyTypeEnumVO string

const (
	// HTTP 回调方式为HTTP
	HTTP NotifyTypeEnumVO = "HTTP"
	// MQ 回调方式为MQ
	MQ NotifyTypeEnumVO = "MQ"
)

// NotifyTypeEnumVOValueOf 将字符串转换为NotifyTypeEnumVO
func NotifyTypeEnumVOValueOf(s string) NotifyTypeEnumVO {
	switch s {
	case "HTTP":
		return HTTP
	case "MQ":
		return MQ
	default:
		return ""
	}
}

// String 返回NotifyTypeEnumVO的字符串表示
func (n NotifyTypeEnumVO) String() string {
	switch n {
	case HTTP:
		return "HTTP"
	case MQ:
		return "MQ"
	default:
		return ""
	}
}

// NotifyConfigVO 回调配置值对象
type NotifyConfigVO struct {
	// 回调方式；MQ、HTTP
	NotifyType NotifyTypeEnumVO `json:"notifyType"`
	// 回调消息
	NotifyMQ string `json:"notifyMQ"`
	// 回调地址
	NotifyUrl string `json:"notifyUrl"`
}
