package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/auth/model"
)

type authRepositoryImpl struct {
	DB *sqlx.DB
}

func (u authRepositoryImpl) CreateUser(um model.User) error {
	stmt, err := u.DB.Preparex(`INSERT INTO users (name, username, password) VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(um.Name, um.Username, um.Password); err != nil {
		return err
	}

	return nil
}

func (u authRepositoryImpl) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	stmt := `SELECT id, name, username, password FROM users WHERE username = $1 LIMIT 1`

	if err := u.DB.Get(&user, stmt, username); err != nil {
		return user, err
	}

	return user, nil
}

func (u authRepositoryImpl) GetUserByID(id uuid.UUID) (model.User, error) {
	var user model.User

	stmt := `SELECT id, name, username, password FROM users WHERE id = $1 LIMIT 1`

	if err := u.DB.Get(&user, stmt, id); err != nil {
		return user, err
	}

	return user, nil
}

func NewAuthRepository(DB *sqlx.DB) *authRepositoryImpl {
	return &authRepositoryImpl{DB: DB}
}
