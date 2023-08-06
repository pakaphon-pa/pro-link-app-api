package utils

import (
	"fmt"
	"net/http"
	"pro-link-api/internal/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type TokenDetail struct {
	AccessToken  string
	AccessExp    int64
	AccessUuid   string
	RefreshToken string
	RefreshExp   int64
	RefreshUuid  string
}

type CustomClaim struct {
	AccessUuid string
	Authorize  bool
	Exp        int64
	Email      string
	jwt.StandardClaims
}

type AccessData struct {
	Email      string
	AccessUuid string
}

func CreateToken(jwtConfig *config.JwtConfig, email string) (TokenDetail, error) {
	accessUuid := uuid.NewV4().String()
	atClamins := CustomClaim{}
	atClamins.Authorize = true
	atClamins.AccessUuid = accessUuid
	atClamins.Email = email
	atClamins.Exp = time.Now().Add(time.Minute * time.Duration(jwtConfig.AccessMaxAge)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClamins)
	accessToken, _ := at.SignedString([]byte(jwtConfig.AccessSecret))

	refreshUuid := uuid.NewV4().String()
	rtClamins := CustomClaim{}
	rtClamins.Authorize = true
	rtClamins.AccessUuid = refreshUuid
	rtClamins.Email = email
	rtClamins.Exp = time.Now().Add(time.Minute * time.Duration(jwtConfig.RefreshMaxAge)).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClamins)
	refreshToekn, _ := rt.SignedString([]byte(jwtConfig.RefreshSecret))

	return TokenDetail{
		AccessToken:  accessToken,
		AccessExp:    atClamins.Exp,
		AccessUuid:   accessUuid,
		RefreshToken: refreshToekn,
		RefreshExp:   rtClamins.Exp,
		RefreshUuid:  refreshUuid,
	}, nil
}

func VerifyToken(jwtConfig *config.JwtConfig, tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.AccessSecret), nil
	})

	return token, err
}

func ExtractMetaData(jwtConfig *config.JwtConfig, tokenString string) (*AccessData, error) {
	token, err := VerifyToken(jwtConfig, tokenString)

	if err != nil {
		return nil, err
	}

	result := AccessData{}
	claim, ok := token.Claims.(CustomClaim)

	if ok && token.Valid {
		result.AccessUuid = claim.AccessUuid
		result.Email = claim.Email
	}

	return &result, err
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
