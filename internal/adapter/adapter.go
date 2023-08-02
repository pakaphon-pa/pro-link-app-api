package adapter

import "pro-link-api/internal/service"

type Adapter struct {
	service     *service.Service
	authService *service.AuthService
}

func New(svr *service.Service) *Adapter {
	return &Adapter{
		service:     svr,
		authService: service.NewAuthService(svr),
	}
}
