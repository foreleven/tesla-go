package core

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthToken struct {
	token   string
	payload TokenPayload
}

type TokenPayload struct {
	AccessToken string
}

type JwtTokenClaims struct {
	AccessToken string
	jwt.RegisteredClaims
}

const (
	DEFAULT_TOKEN_TIMEOUT = 60
)

func Verify(tokenString string) (*AuthToken, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &JwtTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(JWT_SIGN_KEY)), nil
	})
	if err != nil {
		return nil, errors.New("bad token")
	}
	if cliams, ok := jwtToken.Claims.(*JwtTokenClaims); ok && jwtToken.Valid {
		return &AuthToken{
			token: tokenString,
			payload: TokenPayload{
				AccessToken: cliams.AccessToken,
			},
		}, nil
	} else {
		return nil, errors.New("bad token")
	}

}

func Create(accessToken string) (*AuthToken, error) {
	timeout := GetJwtTokenTimeout()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtTokenClaims{
		accessToken,
		jwt.RegisteredClaims{
			Issuer:    "tesla-go",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Second * time.Duration(timeout))},
		},
	})
	jwtKey, _ := GetJwtSignKey()
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return &AuthToken{
			token: ss,
			payload: TokenPayload{
				AccessToken: accessToken,
			},
		}, nil
	}
	return nil, err
}

func (t AuthToken) Token() string {
	return t.token
}

func GetJwtSignKey() (interface{}, error) {
	configKey := os.Getenv(JWT_SIGN_KEY)
	if configKey == "" {
		return []byte(""), nil
	} else {
		return []byte(configKey), nil
	}
}

func GetJwtTokenTimeout() int64 {
	timeout := os.Getenv(JWT_TOKEN_TIMEOUT)
	if timeout != "" {
		i, err := strconv.ParseInt(timeout, 0, 16)
		if err == nil {
			return i
		} else {
			return DEFAULT_TOKEN_TIMEOUT
		}
	}
	return DEFAULT_TOKEN_TIMEOUT
}
