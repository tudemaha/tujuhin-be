package service

import "github.com/tudemaha/tujuhin-be/internal/auth/dto"

type AuthService interface {
	Register(ud dto.UserRegister) error
	Login(ud dto.UserLogin) (dto.Token, error)
	ValidateToken(token string) (string, error)
	GenerateAccessToken(token string) (dto.Token, error)
}
