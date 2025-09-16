package service

import (
	"github.com/tudemaha/tujuhin-be/internal/auth/dto"
	"github.com/tudemaha/tujuhin-be/internal/auth/model"
	"github.com/tudemaha/tujuhin-be/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	authRepository repository.AuthRepository
}

func (s authServiceImpl) Register(ud dto.UserRegister) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(ud.Password), 10)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     ud.Name,
		Username: ud.Username,
		Password: string(hashedPass),
	}

	if err := s.authRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func ProvideUserRepository(ar repository.AuthRepository) *authServiceImpl {
	return &authServiceImpl{authRepository: ar}
}
