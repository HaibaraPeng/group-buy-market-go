package server

import (
	"net/http"

	"group-buy-market-go/internal/application"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/domain/activity/service"
	"group-buy-market-go/internal/infrastructure/dao"
)

type Server struct {
	router          *http.ServeMux
	groupBuyHandler *application.GroupBuyHandler
}

func NewServer(
	activityRepo dao.GroupBuyActivityDAO,
	groupBuyService *domain.GroupBuyService,
	marketService *service.IIndexGroupBuyMarketService,
) *Server {
	// Create handlers
	groupBuyHandler := application.NewGroupBuyHandler(groupBuyService, activityRepo, marketService)

	return &Server{
		router:          http.NewServeMux(),
		groupBuyHandler: groupBuyHandler,
	}
}

func (s *Server) Route(pattern string, handler http.HandlerFunc) {
	s.router.HandleFunc(pattern, handler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// RegisterRoutes registers all HTTP routes
func (s *Server) RegisterRoutes() {
	s.Route("/market/trial", s.groupBuyHandler.MarketTrial)
}
