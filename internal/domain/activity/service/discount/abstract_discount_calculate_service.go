package discount

import (
	"context"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
)

// AbstractDiscountCalculateService 折扣计算服务抽象基类
type AbstractDiscountCalculateService struct {
	doCalculateFunc    func(ctx context.Context, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float
	log                *log.Helper
	activityRepository *repository.ActivityRepository // 添加ActivityRepository依赖
}

// Ensure AbstractDiscountCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*AbstractDiscountCalculateService)(nil)

// SetDoCalculateFunc 设置具体的折扣计算实现函数
func (s *AbstractDiscountCalculateService) SetDoCalculateFunc(f func(ctx context.Context, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float) {
	s.doCalculateFunc = f
}

// SetLogger 设置日志记录器
func (s *AbstractDiscountCalculateService) SetLogger(logger log.Logger) {
	s.log = log.NewHelper(logger)
}

// SetActivityRepository 设置活动仓储服务
func (s *AbstractDiscountCalculateService) SetActivityRepository(activityRepository *repository.ActivityRepository) {
	s.activityRepository = activityRepository
}

// Calculate 折扣计算
func (s *AbstractDiscountCalculateService) Calculate(ctx context.Context, userId string, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	// 1. 人群标签过滤
	if groupBuyDiscount.DiscountType == model.TagDiscount { // TAG 类型
		isCrowdRange := s.filterTagId(userId, groupBuyDiscount.TagId)
		if !isCrowdRange {
			return originalPrice
		}
	}

	// 2. 折扣优惠计算
	return s.doCalculate(ctx, originalPrice, groupBuyDiscount)
}

// filterTagId 人群过滤 - 限定人群优惠
func (s *AbstractDiscountCalculateService) filterTagId(userId string, tagId string) bool {
	// 调用仓储服务的IsTagCrowdRange方法
	if s.activityRepository != nil {
		result, err := s.activityRepository.IsTagCrowdRange(context.Background(), tagId, userId)
		if err != nil {
			// 如果出现错误，默认返回true，避免影响业务
			s.log.Errorw("调用IsTagCrowdRange方法出错", "error", err, "userId", userId, "tagId", tagId)
			return true
		}
		return result
	}
	// 如果没有设置仓储服务，默认返回true
	return true
}

// doCalculate 抽象方法，由子类实现具体的折扣计算逻辑
func (s *AbstractDiscountCalculateService) doCalculate(ctx context.Context, originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	if s.doCalculateFunc != nil {
		return s.doCalculateFunc(ctx, originalPrice, groupBuyDiscount)
	}
	// 默认实现
	return nil
}
