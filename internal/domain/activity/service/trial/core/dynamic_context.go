package core

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"
)

// DynamicContext 动态上下文
// 在策略树执行过程中传递的动态上下文信息
type DynamicContext struct {
	// 拼团活动营销配置值对象
	GroupBuyActivityDiscountVO *model.GroupBuyActivityDiscountVO `json:"-"` // 不序列化到JSON
	// 商品信息
	SkuVO *model.SkuVO `json:"-"` // 不序列化到JSON
	// 折扣价格
	DeductionPrice *big.Float `json:"-"` // 不序列化到JSON
	// 是否可见
	Visible bool `json:"-"`
	// 是否可用
	Enable bool `json:"-"`
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

// SetDeductionPrice 设置折扣价格
func (d *DynamicContext) SetDeductionPrice(price *big.Float) {
	d.DeductionPrice = price
}

// GetDeductionPrice 获取折扣价格
func (d *DynamicContext) GetDeductionPrice() *big.Float {
	return d.DeductionPrice
}

// SetVisible 设置是否可见
func (d *DynamicContext) SetVisible(visible bool) {
	d.Visible = visible
}

// IsVisible 获取是否可见
func (d *DynamicContext) IsVisible() bool {
	return d.Visible
}

// SetEnable 设置是否可用
func (d *DynamicContext) SetEnable(enable bool) {
	d.Enable = enable
}

// IsEnable 获取是否可用
func (d *DynamicContext) IsEnable() bool {
	return d.Enable
}
