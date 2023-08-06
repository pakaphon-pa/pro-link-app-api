package service

import (
	"context"
	"fmt"
	"net/http"
	"pro-link-api/api"
	"pro-link-api/internal/model"
	"pro-link-api/internal/pkg/exceptions"
	utils "pro-link-api/internal/pkg/utils"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Authenication(c context.Context, auth *api.LoginRequest) (*api.AuthenicationResponse, error)
	Register(c context.Context, account *api.RegisterRequest) (*api.AuthenicationResponse, error)
}

func (s *AuthService) Authenication(c context.Context, auth *api.LoginRequest) (*api.AuthenicationResponse, error) {

	account, err := s.AccountStorage.FindByEmailOrName(c, "", auth.Username)

	if err != nil {
		return nil, err
	}

	if account.AccID == 0 {
		return nil, exceptions.NewWithStatus(http.StatusUnauthorized, "Unauthorizated", "Invalid credentials")
	}

	isPass, err := s.comparePassword(c, &auth.Password, account)
	if err != nil {
		return nil, err
	}

	if !isPass {
		return nil, exceptions.NewWithStatus(http.StatusUnauthorized, "Unauthorizated", "Invalid credentials")
	}

	token, err := utils.CreateToken(&s.Config.JwtConfig, account.AccEmail)

	if err != nil {
		return nil, err
	}

	err = s.CreateAuth(uint64(account.AccID), &token)
	if err != nil {
		return nil, err
	}

	return &api.AuthenicationResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *AuthService) Register(c context.Context, account *api.RegisterRequest) (*api.AuthenicationResponse, error) {

	tx, err := utils.GetTrx(c)
	if err != nil {
		return nil, err
	}

	err = s.validateRegister(c, account)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := s.hashedPassword(c, account)

	if err != nil {
		return nil, err
	}

	toModel := &model.Account{
		AccUsername:  account.Username,
		AccPassword:  *hashedPassword,
		AccEmail:     account.Email,
		AccCreatedBy: 1,
	}

	saved, err := s.AccountStorage.Save(tx, c, toModel)
	if err != nil {
		return nil, err
	}

	token, err := utils.CreateToken(&s.Config.JwtConfig, account.Email)
	if err != nil {
		return nil, err
	}

	err = s.CreateAuth(uint64(saved.AccID), &token)
	if err != nil {
		return nil, err
	}

	return &api.AuthenicationResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *AuthService) validateRegister(c context.Context, account *api.RegisterRequest) error {

	duplicateUsername, err := s.AccountStorage.FindByEmailOrName(c, "", account.Username)
	if err != nil {
		return err
	}

	if duplicateUsername.AccID != 0 {
		return exceptions.NewWithStatus(http.StatusBadRequest, "Bad Request", "Username already exists")
	}

	duplicateEmail, err := s.AccountStorage.FindByEmailOrName(c, account.Email, "")
	if err != nil {
		return err
	}

	if duplicateEmail.AccID != 0 {
		return exceptions.NewWithStatus(http.StatusBadRequest, "Bad Request", "Email already exists")
	}

	if account.ConfirmPassword != account.Password {
		return exceptions.NewWithStatus(http.StatusBadRequest, "Bad Request", "Passwords do not match")
	}

	return nil
}

func (s *AuthService) hashedPassword(c context.Context, account *api.RegisterRequest) (*string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	toString := string(hash)
	return &toString, nil
}

func (s *AuthService) comparePassword(c context.Context, password *string, account *model.Account) (bool, error) {
	byteHash := []byte(account.AccPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(*password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *AuthService) CreateAuth(userId uint64, td *utils.TokenDetail) error {
	at := time.Unix(td.AccessExp, 0)
	rt := time.Unix(td.RefreshExp, 0)
	now := time.Now()
	fmt.Println(td.AccessUuid)
	errAccess := s.Storage.GetRedis().Set(td.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefrech := s.Storage.GetRedis().Set(td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub((now))).Err()
	if errRefrech != nil {
		return errRefrech
	}

	return nil
}
