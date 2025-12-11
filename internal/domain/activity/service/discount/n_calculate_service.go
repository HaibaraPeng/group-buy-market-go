package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"log"
	"math/big"
)

// NCalculateService N元购优惠计算服务
type NCalculateService struct {
	*AbstractDiscountCalculateService
}

// Ensure NCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*NCalculateService)(nil)

// NewNCalculateService 创建N元购优惠计算服务实例
func NewNCalculateService() *NCalculateService {
	service := &NCalculateService{
		AbstractDiscountCalculateService: &AbstractDiscountCalculateService{},
	}
	service.SetDoCalculateFunc(service.doCalculate)
	return service
}

// doCalculate 实现N元购优惠计算逻辑
func (s *NCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	log.Printf("优惠策略折扣计算:%v", groupBuyDiscount.DiscountType)

	// 折扣表达式 - 直接为优惠后的金额
	marketExpr := groupBuyDiscount.MarketExpr

	// n元购
	nPrice, _, err := big.ParseFloat(marketExpr, 10, 64, big.ToZero)
	if err != nil {
		log.Printf("解析N元购价格失败: %v", err)
		return originalPrice
	}

	return nPrice
}
