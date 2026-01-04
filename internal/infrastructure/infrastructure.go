package infrastructure

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/infrastructure/adapter/port"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/dcc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	dao.ProviderSet,
	data.ProviderSet,
	dcc.ProviderSet,
	repository.ProviderSet,
	port.NewTradePort,
)
