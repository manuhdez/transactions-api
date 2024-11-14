package http

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// GetUserIdFromContext extracts the user id claim from a jwt token in context
func GetUserIdFromContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("[GetUserIdFromContext][err: token not found in context]")
		}

		// by default claims is of type `jwt.MapClaims`
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("[GetUserIdFromContext][err: failed to cast claims as jwt.MapClaims]")
		}

		// get user key from jwt claims
		userId, ok := claims["user"].(string)
		if !ok {
			return errors.New("[GetUserIdFromContext][err: user id not present in claims map]")
		}

		c.Set("userId", userId)

		return next(c)
	}
}
