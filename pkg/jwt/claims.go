package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessPayload struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

type RefreshPayload struct {
	ID   uuid.UUID `json:"id"`
	IsRT bool      `json:"is_rt"`
	jwt.RegisteredClaims
}
