package server

import (
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/conf"
	"group-buy-market-go/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, marketService *service.ActivityService, tagService *service.TagService, dccService *service.DccService, tradeService *service.TradeService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterActivityHTTPHTTPServer(srv, marketService)
	v1.RegisterTagHTTPHTTPServer(srv, tagService)
	v1.RegisterDccHTTPHTTPServer(srv, dccService)
	v1.RegisterTradeHTTPHTTPServer(srv, tradeService)
	return srv
}
