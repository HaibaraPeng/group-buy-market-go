package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"
)

// IDiscountCalculateService 折扣计算服务
type IDiscountCalculateService interface {
	// Calculate 折扣计算
	//
	// 参数:
	//   userId: 用户ID
	//   originalPrice: 商品原始价格
	//   groupBuyDiscount: 折扣计划配置
	//
	// 返回值:
	//   商品优惠价格
	Calculate(userId string, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float
}
