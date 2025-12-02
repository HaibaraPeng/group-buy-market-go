//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"group-buy-market-go/pkg/server"
)

// ServerSet is a provider set for building the server
var ServerSet = wire.NewSet(server.New)
