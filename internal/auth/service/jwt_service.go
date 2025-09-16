package service

type JWTService interface {
	CreateAccessToken(id, name, username string) (string, error)
	CreateRefreshToken(id string) (string, error)
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, error)
}
