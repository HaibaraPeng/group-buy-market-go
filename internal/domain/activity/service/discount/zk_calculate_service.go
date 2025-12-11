package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"log"
	"math/big"
)

// ZKCalculateService 折扣优惠计算服务
type ZKCalculateService struct {
	*AbstractDiscountCalculateService
}

// Ensure ZKCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*ZKCalculateService)(nil)

// NewZKCalculateService 创建折扣优惠计算服务实例
func NewZKCalculateService() *ZKCalculateService {
	service := &ZKCalculateService{
		AbstractDiscountCalculateService: &AbstractDiscountCalculateService{},
	}
	service.SetDoCalculateFunc(service.doCalculate)
	return service
}

// doCalculate 实现折扣优惠计算逻辑
func (s *ZKCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	log.Printf("优惠策略折扣计算:%v", groupBuyDiscount.DiscountType)

	// 折扣表达式 - 折扣百分比
	marketExpr := groupBuyDiscount.MarketExpr

	// 折扣价格
	discountRate, _, err := big.ParseFloat(marketExpr, 10, 64, big.ToZero)
	if err != nil {
		log.Printf("解析折扣率失败: %v", err)
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
