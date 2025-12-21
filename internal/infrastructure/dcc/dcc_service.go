package dcc

import (
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

// DCCService 动态配置中心服务
type DCCService struct {
	downgradeSwitch string
	cutRange        string
}

// NewDCCService 创建DCC服务实例
func NewDCCService() *DCCService {
	// 从环境变量读取配置，如果不存在则使用默认值
	downgradeSwitch := getEnvWithDefault("DOWNGRADE_SWITCH", "0")
	cutRange := getEnvWithDefault("CUT_RANGE", "100")

	return &DCCService{
		downgradeSwitch: downgradeSwitch,
		cutRange:        cutRange,
	}
}

// IsDowngradeSwitch 判断是否开启降级开关
func (d *DCCService) IsDowngradeSwitch() bool {
	return d.downgradeSwitch == "1"
}

// IsCutRange 判断用户是否在切量范围内
func (d *DCCService) IsCutRange(userId string) (bool, error) {
	cutRange, err := strconv.Atoi(d.cutRange)
	if err != nil {
		return false, fmt.Errorf("invalid cut range value: %s", d.cutRange)
	}

	// 计算用户ID的哈希值
	h := fnv.New32a()
	h.Write([]byte(userId))
	hashCode := int(h.Sum32())

	// 获取绝对值
	if hashCode < 0 {
		hashCode = -hashCode
	}

	// 获取最后两位数字
	lastTwoDigits := hashCode % 100

	// 判断是否在切量范围内
	return lastTwoDigits <= cutRange, nil
}

// getEnvWithDefault 获取环境变量，如果不存在则返回默认值
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		// 尝试使用带前缀的环境变量名
		prefixedKey := "GROUPBUY_" + key
		value = os.Getenv(prefixedKey)
	}

	if value == "" {
		return defaultValue
	}
	return value
}

// UpdateConfig 更新配置值
func (d *DCCService) UpdateConfig(key, value string) {
	switch strings.ToLower(key) {
	case "downgradeswitch":
		d.downgradeSwitch = value
	case "cutrange":
		d.cutRange = value
	}
}
