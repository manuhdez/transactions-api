package infra

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenDuration = time.Hour * 24

type TokenService interface {
	CreateToken(userId string) (string, error)
	ValidateToken(token string) bool
}

type JWTService struct {
	secret     string
	expiration time.Time
}

func NewJWTService() JWTService {
	secret := os.Getenv("JWT_SECRET")

	return JWTService{
		secret:     secret,
		expiration: time.Now().Add(TokenDuration),
	}
}

func (t JWTService) Expiration() time.Time {
	return t.expiration
}

func (t JWTService) CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"exp":  t.expiration.Unix(),
		"user": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.secret))
}

func (t JWTService) ValidateToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if err != nil {
		log.Printf("could not parse token: %e", err)
		return false
	}

	return token.Valid
}
