package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"
)

// AbstractDiscountCalculateService 折扣计算服务抽象类
type AbstractDiscountCalculateService struct {
	doCalculateFunc func(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float
}

// SetDoCalculateFunc 设置折扣计算逻辑
func (s *AbstractDiscountCalculateService) SetDoCalculateFunc(f func(roriginalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float) {
	s.doCalculateFunc = f
}

// Calculate 折扣计算
func (s *AbstractDiscountCalculateService) Calculate(userId string, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	// 1. 人群标签过滤
	// 注意：这里假设 DiscountType 为 byte 类型，其中值为 2 表示 TAG 类型
	// 在实际项目中，应该定义常量来表示这个值
	if groupBuyDiscount.DiscountType == 2 { // TAG 类型
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
	if s.doCalculateFunc != nil {
		return s.doCalculateFunc(originalPrice, groupBuyDiscount)
	}
	// 默认实现
	return nil
}
