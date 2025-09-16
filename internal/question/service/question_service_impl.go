package service

import (
	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
	"github.com/tudemaha/tujuhin-be/internal/question/repository"
)

type questionServiceImpl struct {
	questionRepo repository.QuestionRepository
}

func (s questionServiceImpl) CreateQuestion(qd dto.QuestionRequestBody, owner uuid.UUID) error {
	question := model.Question{
		UserID:   owner,
		Question: qd.Question,
	}

	if err := s.questionRepo.CreateQuestion(question); err != nil {
		return err
	}
	return nil
}

func ProvideQuestionService(qr repository.QuestionRepository) *questionServiceImpl {
	return &questionServiceImpl{questionRepo: qr}
}
