//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"group-buy-market-go/pkg/server"
)

func initializeServer() (*server.Server, error) {
	panic(wire.Build(
		ServerSet,
	))
}
