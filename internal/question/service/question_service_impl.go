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

func (s questionServiceImpl) CreateQuestion(qd dto.QuestionRequestBody, owner string) error {
	userID := uuid.MustParse(owner)
	question := model.Question{
		UserID:   userID,
		Question: qd.Question,
	}

	if err := s.questionRepo.CreateQuestion(question); err != nil {
		return err
	}
	return nil
}

func NewQuestionService(qr repository.QuestionRepository) *questionServiceImpl {
	return &questionServiceImpl{questionRepo: qr}
}
