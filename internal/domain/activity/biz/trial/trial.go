package trial

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/domain/activity/biz/trial/node"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	node.NewRootNode,
	node.NewEndNode,
	node.NewErrorNode,
	node.NewMarketNode,
	node.NewSwitchNode,
	node.NewTagNode,
)
