package repository

import (
	"github.com/google/wire"
)

// ProviderSet for wire
var ProviderSet = wire.NewSet(
	NewActivityRepository,
	NewTagRepository,
	NewTradeRepository,
)
