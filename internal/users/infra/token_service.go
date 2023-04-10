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

type TokenService interface {
	CreateToken(userId string) (string, error)
	ValidateToken(token string) bool
}

type JWTService struct {
	expiration time.Time
}

func NewJWTService() JWTService {
	return JWTService{
		expiration: time.Now().Add(tokenDuration),
	}
}

func (t JWTService) CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  t.expiration,
		"user": userId,
	})

	return token.SignedString([]byte(jwtSecretKey))
}

func (t JWTService) ValidateToken(token string) bool {
	return false
}
