//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure"
	httpInterface "group-buy-market-go/internal/interfaces/http"
)

func initializeServer(db *sql.DB) (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
		infrastructure.NewMySQLGroupBuyActivityRepository,
		domain.NewGroupBuyService,
		wire.Bind(new(domain.GroupBuyActivityRepository), new(*infrastructure.MySQLGroupBuyActivityRepository)),
	))
}
