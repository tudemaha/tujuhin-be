package service

import (
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
)

type QuestionService interface {
	CreateQuestion(qd dto.QuestionRequestBody, owner string) error
}
