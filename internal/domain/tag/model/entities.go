package model

import "time"

// CrowdTagsJobEntity 批次任务对象
type CrowdTagsJobEntity struct {
	// 标签类型（参与量、消费金额）
	TagType int `json:"tagType"`
	// 标签规则（限定类型 N次）
	TagRule string `json:"tagRule"`
	// 统计数据，开始时间
	StatStartTime time.Time `json:"statStartTime"`
	// 统计数据，结束时间
	StatEndTime time.Time `json:"statEndTime"`
}
