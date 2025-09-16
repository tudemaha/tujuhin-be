package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/auth/dto"
	"github.com/tudemaha/tujuhin-be/internal/auth/model"
	"github.com/tudemaha/tujuhin-be/internal/auth/repository"
)

type authServiceImpl struct {
	authRepository repository.AuthRepository
	hasher         HashService
	jwt            JWTService
}

func (s authServiceImpl) Register(ud dto.UserRegister) error {
	hashedPass, err := s.hasher.Hash(ud.Password)
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

func (s authServiceImpl) Login(ud dto.UserLogin) (dto.Token, error) {
	var tokens dto.Token

	user, err := s.authRepository.GetUserByUsername(ud.Username)
	if err != nil {
		return tokens, nil
	}

	err = s.hasher.Compare(user.Password, ud.Password)
	if err != nil {
		return tokens, err
	}

	tokens.AccessToken, err = s.jwt.CreateAccessToken(user.ID.String(), user.Name, user.Username)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken, err = s.jwt.CreateRefreshToken(user.ID.String())
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func (s authServiceImpl) ValidateToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("invalid access token")
	}

	id, err := s.jwt.Validate(token)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s authServiceImpl) GenerateAccessToken(token string) (dto.Token, error) {
	var tokens dto.Token
	if token == "" {
		return tokens, errors.New("invalid refresh token")
	}

	id, err := s.jwt.Validate(token)
	if err != nil {
		return tokens, err
	}

	user, err := s.authRepository.GetUserByID(uuid.MustParse(id))
	if err != nil {
		return tokens, err
	}

	tokens.AccessToken, err = s.jwt.CreateAccessToken(user.ID.String(), user.Name, user.Username)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken, err = s.jwt.CreateRefreshToken(user.ID.String())
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func NewAuthService(ar repository.AuthRepository, passHasher HashService, jwt JWTService) *authServiceImpl {
	return &authServiceImpl{authRepository: ar, hasher: passHasher, jwt: jwt}
}
