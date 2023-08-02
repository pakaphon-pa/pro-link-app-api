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
	RefreshToekn string
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

func CreateToken(jwtConfig *config.JwtConfig) (TokenDetail, error) {
	accessUuid := uuid.NewV4().String()

	atClamins := CustomClaim{}
	atClamins.Authorize = true
	atClamins.AccessUuid = accessUuid
	atClamins.Exp = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClamins)
	aToken, _ := at.SignedString([]byte(jwtConfig.AccessSecret))
	fmt.Println(aToken)

	return TokenDetail{}, nil
}

func VerifyToken(jwtConfig *config.JwtConfig, r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.AccessSecret), nil
	})

	return token, err
}

func ExtractMetaData(jwtConfig *config.JwtConfig, r *http.Request) (*AccessData, error) {
	token, err := VerifyToken(jwtConfig, r)

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
