//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/application"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure"
	"group-buy-market-go/internal/interfaces/http"
)

// ServerSet is a provider set for building the server
var ServerSet = wire.NewSet(
	http.NewServer,
	application.NewService,
	infrastructure.NewMySQLProductRepository,
	domain.NewProductService,
	infrastructure.NewMySQLGroupBuyActivityRepository,
	domain.NewGroupBuyService,
)
