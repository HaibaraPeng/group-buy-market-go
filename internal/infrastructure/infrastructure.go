package infrastructure

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"group-buy-market-go/internal/infrastructure/cache"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/dcc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	cache.ProviderSet,
	dao.ProviderSet,
	data.ProviderSet,
	dcc.ProviderSet,
	repository.ProviderSet,
)
