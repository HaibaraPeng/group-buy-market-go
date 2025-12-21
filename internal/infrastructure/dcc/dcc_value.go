package dcc

// DCCValue 动态配置中心注解标记
// 在Go中通过结构体标签实现类似Java注解的功能
type DCCValue struct {
	Value string `json:"value"`
}

// NewDCCValue 创建DCCValue实例
func NewDCCValue(value string) *DCCValue {
	return &DCCValue{
		Value: value,
	}
}
