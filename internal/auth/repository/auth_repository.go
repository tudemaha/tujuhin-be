package repository

import "github.com/tudemaha/tujuhin-be/internal/auth/model"

type AuthRepository interface {
	CreateUser(um model.User) error
}
