package activity

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain/activity/biz/discount"
	"group-buy-market-go/internal/domain/activity/biz/trial"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	discount.ProviderSet,
	trial.ProviderSet,
)
