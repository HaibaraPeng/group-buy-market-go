package infrastructure

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	repository.ProviderSet,
)
