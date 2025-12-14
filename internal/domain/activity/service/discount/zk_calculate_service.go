package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
)

// ZKCalculateService 折扣优惠计算服务
type ZKCalculateService struct {
	*AbstractDiscountCalculateService
}

// Ensure ZKCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*ZKCalculateService)(nil)

// NewZKCalculateService 创建折扣优惠计算服务实例
func NewZKCalculateService(logger log.Logger) *ZKCalculateService {
	service := &ZKCalculateService{
		AbstractDiscountCalculateService: &AbstractDiscountCalculateService{},
	}
	service.SetDoCalculateFunc(service.doCalculate)
	service.SetLogger(logger)
	return service
}

// doCalculate 实现折扣优惠计算逻辑
func (s *ZKCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	s.logger.Log(log.LevelInfo, "msg", "优惠策略折扣计算", "discountType", groupBuyDiscount.DiscountType)

	// 折扣表达式 - 折扣百分比
	marketExpr := groupBuyDiscount.MarketExpr

	// 折扣价格
	discountRate, _, err := big.ParseFloat(marketExpr, 10, 64, big.ToZero)
	if err != nil {
		s.logger.Log(log.LevelError, "msg", "解析折扣率失败", "error", err)
		return originalPrice
	}

	deductionPrice := new(big.Float).Mul(originalPrice, discountRate)

	// 判断折扣后金额，最低支付1分钱
	zero := big.NewFloat(0)
	minPrice := big.NewFloat(0.01)
	if deductionPrice.Cmp(zero) <= 0 {
		return minPrice
	}

	return deductionPrice
}
