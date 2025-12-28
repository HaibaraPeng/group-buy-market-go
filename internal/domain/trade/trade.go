package trade

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain/trade/biz/filter"
)

// ProviderSet for wire
var ProviderSet = wire.NewSet(
	filter.NewActivityUsabilityRuleFilter,
	filter.NewUserTakeLimitRuleFilter,
	filter.NewTradeRuleFilterFactory,
)
