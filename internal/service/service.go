package service

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewActivityService, NewDccService, NewTagService, NewTradeService)
