//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure/dao"
	httpInterface "group-buy-market-go/internal/interfaces/http"
)

func initializeServer(db *gorm.DB) (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
		dao.NewMySQLGroupBuyActivityDAO,
		domain.NewGroupBuyService,
		wire.Bind(new(domain.GroupBuyActivityRepository), new(*dao.MySQLGroupBuyActivityDAO)),
	))
}
