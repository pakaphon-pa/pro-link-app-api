package service

import (
	"context"
	"pro-link-api/api"
)

type IUserService interface {
	SaveProfile(c context.Context, auth *api.ProfileRequest) (*api.SaveResponse, error)
}

func (s *UserService) SaveProfile(c context.Context, auth *api.ProfileRequest) (*api.SaveResponse, error) {

	return &api.SaveResponse{
		Message: "success",
		Code:    "200",
	}, nil
}
