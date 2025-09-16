package repository

import (
	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/auth/model"
)

type AuthRepository interface {
	CreateUser(um model.User) error
	GetUserByUsername(username string) (model.User, error)
	GetUserByID(id uuid.UUID) (model.User, error)
}
