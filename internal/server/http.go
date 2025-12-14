package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/domain/activity/service"
	"group-buy-market-go/internal/infrastructure/dao"
)

type Server struct {
	marketService *service.IIndexGroupBuyMarketService
}

func NewServer(
	activityRepo dao.GroupBuyActivityDAO,
	marketService *service.IIndexGroupBuyMarketService,
) *Server {
	return &Server{
		marketService: marketService,
	}
}

func (s *Server) RegisterHTTPEndpoints(httpSrv *http.Server) {
	v1.RegisterActivityHTTPHTTPServer(httpSrv, s.marketService)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// RegisterRoutes registers all HTTP routes
func (s *Server) RegisterRoutes() {
	s.Route("/market/trial", s.groupBuyHandler.MarketTrial)
}
