//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	httpInterface "group-buy-market-go/internal/interfaces/http"
)

func initializeServer() (*httpInterface.Server, error) {
	panic(wire.Build(
		httpInterface.NewServer,
	))
}
