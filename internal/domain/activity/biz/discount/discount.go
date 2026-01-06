package discount

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewZJCalculateService,
	NewZKCalculateService,
	NewMJCalculateService,
	NewNCalculateService,
)
