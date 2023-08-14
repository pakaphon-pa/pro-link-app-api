package service

import (
	"context"
	"fmt"
	"net/http"
	"pro-link-api/api"
	"pro-link-api/internal/client"
	"pro-link-api/internal/model"
	"pro-link-api/internal/pkg/exceptions"
	utils "pro-link-api/internal/pkg/utils"
	"strconv"
	"time"

	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Authenication(c context.Context, auth *api.LoginRequest) (*api.AuthenicationResponse, error)
	Register(c context.Context, account *api.RegisterRequest) (*api.AuthenicationResponse, error)
	SendVerifyAccountEmail(c context.Context) (*api.SaveResponse, error)
	VerifyEmail(c context.Context, verificationCode string) (*api.SaveResponse, error)
	Me(c context.Context) (*api.ProfileResponse, error)
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

	code := uuid.NewV4().String()
	verificationCode := utils.Encode(code)

	toModel := &model.Account{
		AccUsername:         account.Username,
		AccPassword:         *hashedPassword,
		AccEmail:            account.Email,
		AccVerificationCode: &code,
		AccCreatedBy:        1,
	}

	saved, err := s.AccountStorage.Save(tx, c, toModel)
	if err != nil {
		return nil, err
	}

	emailReq := &client.VerifyEamilRequest{
		Email:      saved.AccEmail,
		VerifyCode: verificationCode,
		Name:       saved.AccUsername,
		Url:        s.Config.Server.ClientOrigin + "/email/" + "verify/" + verificationCode,
	}

	err = s.NotificationClient.SendVerifyAccountEmail(emailReq)
	if err != nil {
		fmt.Println("Error sending verify account email")
		return nil, err
	}
	fmt.Println("Error sending verify account email")

	token, err := utils.CreateToken(&s.Config.JwtConfig, account.Email)
	if err != nil {

		return nil, err
	}

	err = s.CreateAuth(uint64(saved.AccID), &token)
	if err != nil {
		return nil, err
	}

	fmt.Println("Create Account Successfully")
	return &api.AuthenicationResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *AuthService) SendVerifyAccountEmail(c context.Context) (*api.SaveResponse, error) {

	tx, id, _, err := utils.GetUserIdAndTrx(c)
	if err != nil {
		return nil, err
	}

	account, err := s.AccountStorage.FindById(c, id)

	if err != nil {
		return nil, exceptions.NewWithStatus(http.StatusBadRequest, "Not Found", "Account not found")
	}

	if account.AccIsVerified {
		return nil, exceptions.NewWithStatus(http.StatusConflict, "Fail", "account already verified")

	}

	code := uuid.NewV4().String()
	verificationCode := utils.Encode(code)

	account.AccVerificationCode = &code

	_, err = s.AccountStorage.Save(tx, c, account)
	if err != nil {
		return nil, err
	}

	emailReq := &client.VerifyEamilRequest{
		Email:      account.AccEmail,
		VerifyCode: verificationCode,
		Name:       account.AccUsername,
		Url:        s.Config.Server.ClientOrigin + "/email/" + "verify/" + verificationCode,
	}

	err = s.NotificationClient.SendVerifyAccountEmail(emailReq)
	if err != nil {
		fmt.Println("Error sending verify account email")
		return nil, err
	}
	fmt.Println("Error sending verify account email")

	if err != nil {
		return nil, err
	}

	return &api.SaveResponse{
		Message: "Email sent",
		Code:    "Success",
	}, nil
}

func (s *AuthService) Me(c context.Context) (*api.ProfileResponse, error) {

	userId, err := utils.GetUserId(c)

	if err != nil {
		return nil, err
	}

	result := &api.ProfileResponse{
		AccId: userId,
	}

	account, err := s.AccountStorage.FindById(c, userId)

	if err != nil {
		return nil, err
	}

	result.AccId = account.AccID

	profile, err := s.ProfileStorage.FindByAccId(c, userId)

	if err != nil {
		return nil, err
	}

	result.FirstName = profile.PrfFirstName
	result.LastName = profile.PrfLastName
	result.About = profile.PrfAbout
	result.Address = profile.PrfAddress
	result.PhoneNumber = profile.PrfPhoneNumber
	result.PhoneType = profile.PrfPhoneType
	result.Address = profile.PrfAddress

	edu, err := s.EducationStorage.FindByAccId(c, userId)

	if err != nil {
		return nil, err
	}

	result.Education = model.ToEducationListDoamin(edu)

	exp, err := s.ExperienceStorage.FindByAccId(c, userId)

	if err != nil {
		return nil, err
	}

	result.Experience = model.ToExperienceListDoamin(exp)

	skill, err := s.SkillStorage.FindByAccId(c, userId)

	if err != nil {
		return nil, err
	}

	result.Skill = model.ToSkillListDoamin(skill)

	lan, err := s.LanguageStorage.FindByAccId(c, userId)

	if err != nil {
		return nil, err
	}

	result.Language = model.ToLanguageListDoamin(lan)

	return result, nil
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

func (s *AuthService) VerifyEmail(c context.Context, verificationCode string) (*api.SaveResponse, error) {
	fmt.Println("Verifying email ")

	tx, err := utils.GetTrx(c)
	if err != nil {
		return nil, err
	}

	current := time.Now()

	decode, err := utils.Decode(verificationCode)
	if err != nil {
		return nil, err
	}

	result, err := s.AccountStorage.FindByVerificationCode(c, decode)

	if err != nil || result.AccID == 0 {
		return nil, exceptions.NewWithStatus(http.StatusBadRequest, "Not Found", "Invalid verification code")
	}

	if result.AccIsVerified {
		return nil, exceptions.NewWithStatus(http.StatusConflict, "Fail", "account already verified")

	}

	result.AccVerificationCode = nil
	result.AccIsVerified = true
	result.AccUpdatedDate = &current

	_, err = s.AccountStorage.Save(tx, c, result)

	if err != nil {
		return nil, err
	}

	return &api.SaveResponse{
		Message: "Email verified",
		Code:    "Success",
	}, nil
}
