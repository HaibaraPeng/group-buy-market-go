//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure"
	httpInterface "group-buy-market-go/internal/interfaces/http"
	"gorm.io/gorm"
)

func initializeServer(db *gorm.DB) (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
		infrastructure.NewMySQLProductRepository,
		infrastructure.NewMySQLGroupBuyActivityRepository,
		domain.NewProductService,
		domain.NewGroupBuyService,
		wire.Bind(new(domain.ProductRepository), new(*infrastructure.MySQLProductRepository)),
		wire.Bind(new(domain.GroupBuyActivityRepository), new(*infrastructure.MySQLGroupBuyActivityRepository)),
	))
}