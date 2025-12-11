package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"
)

// AbstractDiscountCalculateService 折扣计算服务抽象类
type AbstractDiscountCalculateService struct{}

// Calculate 折扣计算
func (s *AbstractDiscountCalculateService) Calculate(userId string, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	// 1. 人群标签过滤
	if groupBuyDiscount.DiscountType == model.TagDiscount { // TAG 类型
		isCrowdRange := s.filterTagId(userId, groupBuyDiscount.TagId)
		if !isCrowdRange {
			return originalPrice
		}
	}

	// 2. 折扣优惠计算
	return s.doCalculate(originalPrice, groupBuyDiscount)
}

// filterTagId 人群过滤 - 限定人群优惠
func (s *AbstractDiscountCalculateService) filterTagId(userId string, tagId string) bool {
	// todo xiaofuge 后续开发这部分
	return true
}

// doCalculate 抽象方法，由子类实现具体的折扣计算逻辑
func (s *AbstractDiscountCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	// 子类需要实现此方法
	return nil
}
