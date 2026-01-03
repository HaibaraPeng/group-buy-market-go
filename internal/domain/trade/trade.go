package trade

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain/trade/biz/lock"
	"group-buy-market-go/internal/domain/trade/biz/lock/filter"
	"group-buy-market-go/internal/domain/trade/biz/settlement"
)

// ProviderSet for wire
var ProviderSet = wire.NewSet(
	lock.NewTradeLockOrder,
	settlement.NewTradeSettlementOrderService,
	filter.NewActivityUsabilityRuleFilter,
	filter.NewUserTakeLimitRuleFilter,
	filter.NewTradeLockRuleFilterFactory,
)
