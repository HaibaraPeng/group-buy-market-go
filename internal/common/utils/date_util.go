package utils

import (
	"fmt"
	"time"
)

// DifferenceDateTime2Str 计算两个时间之间的差值并格式化为 HH:MM:SS 格式
func DifferenceDateTime2Str(startTime time.Time, endTime time.Time) string {
	// 检查时间是否有效
	if startTime.IsZero() || endTime.IsZero() {
		return "无效的时间"
	}

	// 计算时间差
	diff := endTime.Sub(startTime)

	// 如果时间差为负数，表示已经结束
	if diff < 0 {
		return "已结束"
	}

	// 获取总秒数
	totalSeconds := int64(diff.Seconds())

	// 计算小时、分钟和秒
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	// 格式化为 HH:MM:SS
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
