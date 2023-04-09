package infra

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecretKey  = os.Getenv("JWT_SECRET")
	tokenDuration = time.Hour * 24
)

type TokenService struct {
	expiration time.Time
}

func NewTokenService() TokenService {
	return TokenService{
		expiration: time.Now().Add(tokenDuration),
	}
}

func (t TokenService) CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  t.expiration,
		"user": userId,
	})

	return token.SignedString([]byte(jwtSecretKey))
}

func (t TokenService) ValidateToken(token string) bool {
	return false
}
