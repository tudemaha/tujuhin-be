package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/global/dto"
)

type QuestionResponse struct {
	ID        uuid.UUID `json:"id"`
	Question  string    `json:"question"`
	TotalVote int       `json:"total_vote"`
	VoteState string    `json:"vote_state"`
	Owner     dto.User  `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
}

type QuestionsResponse []QuestionResponse
