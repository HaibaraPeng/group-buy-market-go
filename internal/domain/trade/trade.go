package trade

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain/trade/biz"
	"group-buy-market-go/internal/domain/trade/biz/lock/filter"
)

// ProviderSet for wire
var ProviderSet = wire.NewSet(
	biz.NewTradeOrder,
	filter.NewActivityUsabilityRuleFilter,
	filter.NewUserTakeLimitRuleFilter,
	filter.NewTradeRuleFilterFactory,
)
