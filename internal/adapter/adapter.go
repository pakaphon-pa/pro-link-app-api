package adapter

import (
	"pro-link-api/internal/service"
)

type Adapter struct {
	service     *service.Service
	authService service.IAuthService
	userService service.IUserService
}

func New(svr *service.Service) *Adapter {
	return &Adapter{
		service:     svr,
		authService: service.NewAuthService(svr),
		userService: service.NewUserService(svr),
	}
}
