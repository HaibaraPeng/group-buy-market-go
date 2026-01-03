package trade

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain/trade/biz/lock"
	lockfilter "group-buy-market-go/internal/domain/trade/biz/lock/filter"
	"group-buy-market-go/internal/domain/trade/biz/settlement"
	settlementfilter "group-buy-market-go/internal/domain/trade/biz/settlement/filter"
)

// ProviderSet for wire
var ProviderSet = wire.NewSet(
	lock.NewTradeLockOrder,
	settlement.NewTradeSettlementOrderService,
	lockfilter.NewActivityUsabilityRuleFilter,
	lockfilter.NewUserTakeLimitRuleFilter,
	lockfilter.NewTradeLockRuleFilterFactory,
	settlementfilter.NewSCRuleFilter,
	settlementfilter.NewSettableRuleFilter,
	settlementfilter.NewOutTradeNoRuleFilter,
	settlementfilter.NewEndRuleFilter,
	settlementfilter.NewTradeSettlementRuleFilterFactory,
)
