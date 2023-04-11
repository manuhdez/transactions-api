package infra_test

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

const tokenSecret = "test-secret"

func TestJWTService_CreateToken(t *testing.T) {
	t.Setenv("JWT_SECRET", tokenSecret)

	userId := "1234-5678"
	service := infra.NewJWTService()

	tokenStr, err := service.CreateToken(userId)
	if err != nil {
		t.Errorf("could not generate a valid token: %e", err)
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		t.Fatalf("could not parse token: %e", err)
	}

	if !token.Valid {
		t.Fatalf("token is not valid")
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["user"] != userId {
		t.Errorf("userId claims got %s, want %s", claims["userId"], userId)
	}

	exp := int64(claims["exp"].(float64))
	wanted := service.Expiration().Unix()
	if exp != wanted {
		t.Errorf("token exp got %v want %v", exp, wanted)
	}
}

func TestJWTService_ValidateToken(t *testing.T) {
	t.Setenv("JWT_SECRET", tokenSecret)
	service := infra.NewJWTService()

	invalidToken := "this-is-an-invalid-token"
	validToken, err := service.CreateToken("test-user-id")
	if err != nil {
		t.Fatalf("error generating a test token: %e", err)
	}

	t.Run("with a valid token", func(t *testing.T) {
		got := service.ValidateToken(validToken)
		if got != true {
			t.Errorf("Validate(token): got %v, want %v", got, true)
		}
	})

	t.Run("with an invalid token", func(t *testing.T) {
		got := service.ValidateToken(invalidToken)
		if got != false {
			t.Errorf("Validate(token): got %v, want %v", got, false)
		}
	})

}
