package http

import (
	"group-buy-market-go/internal/application"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure/dao"
	"net/http"
)

type Server struct {
	router          *http.ServeMux
	groupBuyHandler *application.GroupBuyHandler
}

func NewServer(
	activityRepo dao.GroupBuyActivityDAO,
	groupBuyService *domain.GroupBuyService,
) *Server {
	// Create handlers
	groupBuyHandler := application.NewGroupBuyHandler(groupBuyService, activityRepo)

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
	s.Route("/", s.handleHome)
	s.Route("/test", s.handleTest)
	s.Route("/groupbuy/activity", s.groupBuyHandler.GetActivity)
}

// handleHome handles requests to the home page
func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Group Buy Market!"))
}

// handleTest handles requests to the test endpoint
func (s *Server) handleTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test endpoint is working!"))
}
