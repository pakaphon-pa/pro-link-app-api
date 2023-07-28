package adapter

import "pro-link-api/internal/service"

type Adapter struct {
	service *service.Service
}

func New(service *service.Service) *Adapter {
	return &Adapter{
		service: service,
	}
}
