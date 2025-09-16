package service

import "github.com/tudemaha/tujuhin-be/internal/auth/dto"

type AuthService interface {
	Register(ud dto.UserRegister) error
}
