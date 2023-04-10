package mocks

type TokenService struct {
	Token               string
	CreateTokenError    error
	ValidateTokenResult bool
}

func NewTokenService(token string) TokenService {
	return TokenService{
		Token:               token,
		CreateTokenError:    nil,
		ValidateTokenResult: true,
	}
}

func (m TokenService) SetCreateTokenError(err error) {
	m.CreateTokenError = err
}

func (m TokenService) CreateToken(_ string) (string, error) {
	return m.Token, m.CreateTokenError
}

func (m TokenService) ValidateToken(_ string) bool {
	return m.ValidateTokenResult
}
