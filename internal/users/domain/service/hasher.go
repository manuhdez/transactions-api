package service

type HashService interface {
	Hash(value string) (string, error)
	Compare(hashed, ref string) error
}
