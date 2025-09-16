package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/auth/model"
)

type authRepositoryImpl struct {
	DB *sqlx.DB
}

func (u authRepositoryImpl) CreateUser(um model.User) error {
	stmt, err := u.DB.Preparex("INSERT INTO users (name, username, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(um.Name, um.Username, um.Password); err != nil {
		return err
	}

	return nil
}

func ProvideUserRepository(DB *sqlx.DB) *authRepositoryImpl {
	return &authRepositoryImpl{DB: DB}
}
