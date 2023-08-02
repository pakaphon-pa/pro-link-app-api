package service

import utils "pro-link-api/internal/pkg/utils"

func (s *AuthService) Authenication() {
	_, _ = utils.CreateToken(&s.Config.JwtConfig)
}
