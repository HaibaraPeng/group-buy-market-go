package main

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/conf"
	"group-buy-market-go/internal/domain/activity/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
)

func NewHTTPServer(c *conf.Server, marketService *service.IIndexGroupBuyMarketService, logger log.Logger) *http.Server {
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
	return srv
}

func NewHTTPServerWithContext(c *conf.Server, marketService *service.IIndexGroupBuyMarketService, logger log.Logger) *http.Server {
	return NewHTTPServer(c, marketService, logger)
}
