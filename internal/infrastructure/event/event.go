package event

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/infrastructure/event/listener"
)

// ProviderSet is event providers.
var ProviderSet = wire.NewSet(
	listener.NewTeamSuccessEventListener,
	listener.NewEventListenerManager,
)
