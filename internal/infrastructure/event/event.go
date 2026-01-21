package event

import (
	"github.com/google/wire"
	"group-buy-market-go/internal/infrastructure/event/listener"
	"group-buy-market-go/internal/infrastructure/event/publish"
)

// ProviderSet is event providers.
var ProviderSet = wire.NewSet(
	listener.NewTeamSuccessEventListener,
	listener.NewEventListenerManager,
	publish.NewRabbitMQEventPublisher,
)
