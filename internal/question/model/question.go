package model

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Question  string
	TotalVote int
	CreatedAt time.Time
}
