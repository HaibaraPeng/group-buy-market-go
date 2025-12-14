package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"group-buy-market-go/internal/server"
)

func NewHTTPServer(s *server.Server, opts ...http.ServerOption) *http.Server {
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", s)
	return srv
}

func NewHTTPServerWithContext(ctx context.Context, s *server.Server, opts ...http.ServerOption) *http.Server {
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", s)
	return srv
}
