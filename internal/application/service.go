package application

// Service represents the application service layer
type Service struct{}

// NewService creates a new application service
func NewService() *Service {
	return &Service{}
}

// HealthCheck returns the health status of the application
func (s *Service) HealthCheck() string {
	return "OK"
}
