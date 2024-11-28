package domain_service

// TokenService - Describes a service that is able to generate json web tokens and validate them
type TokenService interface {
	CreateToken(userId string) (string, error)
	ValidateToken(token string) bool
}
