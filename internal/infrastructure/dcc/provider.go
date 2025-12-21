package dcc

import "github.com/google/wire"

// ProviderSet is dcc providers.
var ProviderSet = wire.NewSet(NewDCCService)
