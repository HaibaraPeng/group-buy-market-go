//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure"
	httpInterface "group-buy-market-go/internal/interfaces/http"
)

func initializeServer(db *gorm.DB) (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
		infrastructure.NewMySQLGroupBuyActivityRepository,
		domain.NewGroupBuyService,
		wire.Bind(new(domain.GroupBuyActivityRepository), new(*infrastructure.MySQLGroupBuyActivityRepository)),
	))
}
