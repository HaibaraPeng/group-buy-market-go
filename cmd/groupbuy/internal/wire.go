//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/domain/activity/service"
	"group-buy-market-go/internal/domain/activity/service/discount"
	"group-buy-market-go/internal/domain/activity/service/trial/node"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"group-buy-market-go/internal/infrastructure/dao"
	httpInterface "group-buy-market-go/internal/interfaces/http"
)

func initializeServer(db *gorm.DB) (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
		dao.NewMySQLGroupBuyActivityDAO,
		dao.NewMySQLGroupBuyDiscountDAO,
		dao.NewMySQLSkuDAO,
		repository.NewActivityRepository,
		domain.NewGroupBuyService,
		service.NewIIndexGroupBuyMarketService,
		node.NewEndNode,
		node.NewMarketNode,
		node.NewSwitchNode,
		node.NewRootNode,
		discount.NewZJCalculateService,
		discount.NewZKCalculateService,
		discount.NewMJCalculateService,
		discount.NewNCalculateService,
		wire.Bind(new(dao.GroupBuyActivityDAO), new(*dao.MySQLGroupBuyActivityDAO)),
		wire.Bind(new(dao.GroupBuyDiscountDAO), new(*dao.MySQLGroupBuyDiscountDAO)),
		wire.Bind(new(dao.SkuDAO), new(*dao.MySQLSkuDAO)),
	))
}
