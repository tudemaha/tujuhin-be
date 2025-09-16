package service

type JWTService interface {
	CreateAccessToken(id, name, username string) (string, error)
	CreateRefreshToken(id string) (string, error)
	Validate(token string) (string, error)
}
