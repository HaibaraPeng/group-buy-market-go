package model

// GroupBuyProgressVO 拼团进度值对象
type GroupBuyProgressVO struct {
	CompleteCount int `json:"completeCount"`
	TargetCount   int `json:"targetCount"`
	LockCount     int `json:"lockCount"`
}
