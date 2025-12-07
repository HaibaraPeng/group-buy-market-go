package core

import (
	"group-buy-market-go/internal/domain/activity/model"
)

// DynamicContext 动态上下文
// 在策略树执行过程中传递的动态上下文信息
type DynamicContext struct {
	// 拼团活动营销配置值对象
	GroupBuyActivityDiscountVO *model.GroupBuyActivityDiscountVO `json:"-"` // 不序列化到JSON
	// 商品信息
	SkuVO *model.SkuVO `json:"-"` // 不序列化到JSON
}

// SetGroupBuyActivityDiscountVO 设置拼团活动营销配置值对象
func (d *DynamicContext) SetGroupBuyActivityDiscountVO(vo *model.GroupBuyActivityDiscountVO) {
	d.GroupBuyActivityDiscountVO = vo
}

// GetGroupBuyActivityDiscountVO 获取拼团活动营销配置值对象
func (d *DynamicContext) GetGroupBuyActivityDiscountVO() *model.GroupBuyActivityDiscountVO {
	return d.GroupBuyActivityDiscountVO
}

// SetSkuVO 设置商品信息
func (d *DynamicContext) SetSkuVO(vo *model.SkuVO) {
	d.SkuVO = vo
}

// GetSkuVO 获取商品信息
func (d *DynamicContext) GetSkuVO() *model.SkuVO {
	return d.SkuVO
}
