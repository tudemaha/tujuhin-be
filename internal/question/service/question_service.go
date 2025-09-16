package service

import (
	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
)

type QuestionService interface {
	CreateQuestion(qd dto.QuestionRequestBody, owner uuid.UUID) error
}
