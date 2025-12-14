package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
)

// AbstractDiscountCalculateService 折扣计算服务抽象基类
type AbstractDiscountCalculateService struct {
	doCalculateFunc func(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float
	log             *log.Helper
}

// Ensure AbstractDiscountCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*AbstractDiscountCalculateService)(nil)

// SetDoCalculateFunc 设置具体的折扣计算实现函数
func (s *AbstractDiscountCalculateService) SetDoCalculateFunc(f func(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float) {
	s.doCalculateFunc = f
}

// SetLogger 设置日志记录器
func (s *AbstractDiscountCalculateService) SetLogger(logger log.Logger) {
	s.log = log.NewHelper(logger)
}

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
	if s.doCalculateFunc != nil {
		return s.doCalculateFunc(originalPrice, groupBuyDiscount)
	}
	// 默认实现
	return nil
}
