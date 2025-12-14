package discount

import (
	"group-buy-market-go/common/consts"
	"group-buy-market-go/internal/domain/activity/model"
	"math/big"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

// MJCalculateService 满减优惠计算服务
type MJCalculateService struct {
	*AbstractDiscountCalculateService
}

// Ensure MJCalculateService implements IDiscountCalculateService
var _ IDiscountCalculateService = (*MJCalculateService)(nil)

// NewMJCalculateService 创建满减优惠计算服务实例
func NewMJCalculateService(logger log.Logger) *MJCalculateService {
	service := &MJCalculateService{
		AbstractDiscountCalculateService: &AbstractDiscountCalculateService{},
	}
	service.SetDoCalculateFunc(service.doCalculate)
	service.SetLogger(logger)
	return service
}

// doCalculate 实现满减优惠计算逻辑
func (s *MJCalculateService) doCalculate(originalPrice *big.Float, groupBuyDiscount *model.GroupBuyDiscountVO) *big.Float {
	s.log.Infof("优惠策略折扣计算: %v", groupBuyDiscount.DiscountType)

	// 折扣表达式 - 100,10 满100减10元
	marketExpr := groupBuyDiscount.MarketExpr
	parts := strings.Split(marketExpr, consts.SPLIT)

	if len(parts) < 2 {
		s.log.Warnf("无效的满减表达式: %s", marketExpr)
		return originalPrice
	}

	xStr := strings.TrimSpace(parts[0])
	yStr := strings.TrimSpace(parts[1])

	x, _, err1 := big.ParseFloat(xStr, 10, 64, big.ToZero)
	if err1 != nil {
		s.log.Errorf("解析满减条件失败: %v", err1)
		return originalPrice
	}

	y, _, err2 := big.ParseFloat(yStr, 10, 64, big.ToZero)
	if err2 != nil {
		s.log.Errorf("解析减免金额失败: %v", err2)
		return originalPrice
	}

	// 不满足最低满减约束，则按照原价
	if originalPrice.Cmp(x) < 0 {
		return originalPrice
	}

	// 折扣价格
	deductionPrice := new(big.Float).Sub(originalPrice, y)

	// 判断折扣后金额，最低支付1分钱
	zero := big.NewFloat(0)
	minPrice := big.NewFloat(0.01)
	if deductionPrice.Cmp(zero) <= 0 {
		return minPrice
	}

	return deductionPrice
}
