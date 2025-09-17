package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/auth/model"
)

type QuestionModel struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	Question  string    `db:"question"`
	TotalVote int       `db:"total_vote"`
	CreatedAt time.Time `db:"created_at"`
}

type QuestionWithOwner struct {
	QuestionModel
	User model.User   `db:"user"`
	Vote QuestionVote `db:"vote"`
}

type QuestionVote struct {
	ID         uuid.UUID `db:"id"`
	QuestionID uuid.UUID `db:"question_id"`
	VoteState  *string   `db:"vote_state"`
}

type QuestionsWithOwner []QuestionWithOwner
