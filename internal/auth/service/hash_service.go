package service

type HashService interface {
	Hash(password string) (string, error)
	Compare(hashed, plain string) error
}
