package service

import (
	"context"
	"pro-link-api/api"
	utils "pro-link-api/internal/pkg/utils"
	"strconv"
	"time"
)

func (s *AuthService) Authenication(c context.Context) (*api.AuthenicationResponse, error) {
	token, err := utils.CreateToken(&s.Config.JwtConfig, "")

	if err != nil {
		return nil, err
	}

	return &api.AuthenicationResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *AuthService) CreateAuth(userId uint64, td *utils.TokenDetail) error {
	at := time.Unix(td.AccessExp, 0)
	rt := time.Unix(td.RefreshExp, 0)
	now := time.Now()
	errAccess := s.Storage.GetRedis().Set(td.AccessToken, strconv.Itoa(int(userId)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefrech := s.Storage.GetRedis().Set(td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub((now))).Err()
	if errRefrech != nil {
		return errRefrech
	}

	return nil
}
