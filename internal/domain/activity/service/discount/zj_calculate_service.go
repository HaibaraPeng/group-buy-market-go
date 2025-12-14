package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
)

// ZJCalculateService 直减优惠计算服务
type ZJCalculateService struct {
	*AbstractDiscountCalculateService
}

// Ensure ZJCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*ZJCalculateService)(nil)

// NewZJCalculateService 创建直减优惠计算服务实例
func NewZJCalculateService(logger log.Logger) *ZJCalculateService {
	service := &ZJCalculateService{
		AbstractDiscountCalculateService: &AbstractDiscountCalculateService{},
	}
	service.SetDoCalculateFunc(service.doCalculate)
	service.SetLogger(logger)
	return service
}

// doCalculate 实现直减优惠计算逻辑
func (s *ZJCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	s.log.Infof("优惠策略折扣计算: %v", groupBuyDiscount.DiscountType)

	// 折扣表达式 - 直减为扣减金额
	marketExpr := groupBuyDiscount.MarketExpr

	// 折扣价格
	deductionAmount, _, err := big.ParseFloat(marketExpr, 10, 64, big.ToZero)
	if err != nil {
		s.log.Errorf("解析折扣金额失败: %v", err)
		return originalPrice
	}

	deductionPrice := new(big.Float).Sub(originalPrice, deductionAmount)

	// 判断折扣后金额，最低支付1分钱
	zero := big.NewFloat(0)
	minPrice := big.NewFloat(0.01)
	if deductionPrice.Cmp(zero) <= 0 {
		return minPrice
	}

	return deductionPrice
}
