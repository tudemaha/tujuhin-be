package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Username string    `db:"username"`
	Password string    `db:"password"`
}
