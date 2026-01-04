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
	SUCCESS NotifyTaskHTTPEnumVO = "0000"
	// NULL 空值
	NULL NotifyTaskHTTPEnumVO = "1001"
)
