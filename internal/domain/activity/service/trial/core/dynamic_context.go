package core

// DynamicContext 动态上下文
// 在策略树执行过程中传递的动态上下文信息
type DynamicContext struct {
	UserID     int64  `json:"userId"`     // 用户ID
	ActivityID int64  `json:"activityId"` // 活动ID
	Timestamp  int64  `json:"timestamp"`  // 时间戳
	UserLevel  int32  `json:"userLevel"`  // 用户等级
	ClientIP   string `json:"clientIp"`   // 客户端IP
}
