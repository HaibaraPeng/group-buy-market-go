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
		dao.NewMySQLGroupBuyDiscountDAO,
		dao.NewMySQLSkuDAO,
		domain.NewGroupBuyService,
		wire.Bind(new(dao.GroupBuyActivityDAO), new(*dao.MySQLGroupBuyActivityDAO)),
		wire.Bind(new(dao.GroupBuyDiscountDAO), new(*dao.MySQLGroupBuyDiscountDAO)),
		wire.Bind(new(dao.SkuDAO), new(*dao.MySQLSkuDAO)),
	))
}
