package discount

import (
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
)

// NCalculateService N元购优惠计算服务
type NCalculateService struct {
	*AbstractDiscountCalculateService
}

// Ensure NCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*NCalculateService)(nil)

// NewNCalculateService 创建N元购优惠计算服务实例
func NewNCalculateService(logger log.Logger) *NCalculateService {
	service := &NCalculateService{
		AbstractDiscountCalculateService: &AbstractDiscountCalculateService{},
	}
	service.SetDoCalculateFunc(service.doCalculate)
	service.SetLogger(logger)
	return service
}

// doCalculate 实现N元购优惠计算逻辑
func (s *NCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	s.logger.Log(log.LevelInfo, "msg", "优惠策略折扣计算", "discountType", groupBuyDiscount.DiscountType)

	// 折扣表达式 - 直接为优惠后的金额
	marketExpr := groupBuyDiscount.MarketExpr

	// n元购
	nPrice, _, err := big.ParseFloat(marketExpr, 10, 64, big.ToZero)
	if err != nil {
		s.logger.Log(log.LevelError, "msg", "解析N元购价格失败", "error", err)
		return originalPrice
	}

	return nPrice
}
