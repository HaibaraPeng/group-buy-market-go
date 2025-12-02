//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure"
	httpInterface "group-buy-market-go/internal/interfaces/http"
)

func initializeServer() (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
		infrastructure.NewInMemoryGroupBuyActivityRepository,
		domain.NewGroupBuyService,
		wire.Bind(new(domain.GroupBuyActivityRepository), new(*infrastructure.InMemoryGroupBuyActivityRepository)),
	))
}
