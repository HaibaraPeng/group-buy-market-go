package dcc

import (
	"context"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"group-buy-market-go/internal/common/consts"
)

// DCC 动态配置中心服务
type DCC struct {
	redisClient *redis.Client
	configMap   map[string]*ConfigValue
}

// ConfigValue 配置值包装器
type ConfigValue struct {
	key   string
	value string
}

// NewDCC 创建DCC服务实例
func NewDCC(redisClient *redis.Client) *DCC {
	dcc := &DCC{
		redisClient: redisClient,
		configMap:   make(map[string]*ConfigValue),
	}

	// 初始化默认配置
	dcc.initDefaultConfigs()

	// 启动监听
	go dcc.listenForConfigChanges(context.Background())

	return dcc
}

// initDefaultConfigs 初始化默认配置
func (d *DCC) initDefaultConfigs() {
	d.initConfigValue("downgradeSwitch", "0")
	d.initConfigValue("cutRange", "100")
}

// initConfigValue 初始化配置值
func (d *DCC) initConfigValue(key, defaultValue string) {
	configKey := "group_buy_market_dcc_" + key
	// 检查Redis中是否存在该键值
	ctx := context.Background()
	value, err := d.redisClient.Get(ctx, configKey).Result()
	if err != nil || value == "" {
		// 如果检查出现错误或者不存在该键，设置为默认值，使用默认值
		d.configMap[key] = &ConfigValue{
			key:   configKey,
			value: defaultValue,
		}
		return
	} else {
		// 如果Redis中存在该键，获取其值
		value, err = d.redisClient.Get(ctx, configKey).Result()
		if err != nil {
			value = defaultValue
		}
	}

	d.configMap[key] = &ConfigValue{
		key:   configKey,
		value: value,
	}
}

// listenForConfigChanges 监听配置变化
func (d *DCC) listenForConfigChanges(ctx context.Context) {
	pubsub := d.redisClient.Subscribe(ctx, "group_buy_market_dcc")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		d.handleConfigChange(msg.Payload)
	}
}

// handleConfigChange 处理配置变化
func (d *DCC) handleConfigChange(payload string) {
	parts := strings.Split(payload, consts.SPLIT)
	if len(parts) != 2 {
		return
	}

	attribute := parts[0]
	value := parts[1]
	configKey := "group_buy_market_dcc_" + attribute

	// 更新Redis中的值
	ctx := context.Background()
	d.redisClient.Set(ctx, configKey, value, 0)

	// 更新内存中的值
	if config, exists := d.configMap[attribute]; exists {
		config.value = value
	}
}

// PublishConfigChange 发布配置变更消息
func (d *DCC) PublishConfigChange(ctx context.Context, key, value string) error {
	message := key + consts.SPLIT + value
	return d.redisClient.Publish(ctx, "group_buy_market_dcc", message).Err()
}

// GetValue 获取配置值
func (d *DCC) GetValue(key string) string {
	if config, exists := d.configMap[key]; exists {
		return config.value
	}
	return ""
}

// IsDowngradeSwitch 判断是否开启降级开关
func (d *DCC) IsDowngradeSwitch() bool {
	return d.GetValue("downgradeSwitch") == "1"
}

// IsCutRange 判断用户是否在切量范围内
func (d *DCC) IsCutRange(userId string) (bool, error) {
	cutRangeStr := d.GetValue("cutRange")
	cutRange, err := strconv.Atoi(cutRangeStr)
	if err != nil {
		return false, fmt.Errorf("invalid cut range value: %s", cutRangeStr)
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
