package service

import (
	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
	"github.com/tudemaha/tujuhin-be/internal/question/query"
	"github.com/tudemaha/tujuhin-be/internal/question/repository"
)

type questionServiceImpl struct {
	questionRepo  repository.QuestionRepository
	questionQuery query.QuestionQuery
}

func (s questionServiceImpl) CreateQuestion(qd dto.QuestionRequestBody, owner string) error {
	userID := uuid.MustParse(owner)
	question := model.QuestionModel{
		UserID:   userID,
		Question: qd.Question,
	}

	if err := s.questionRepo.CreateQuestion(question); err != nil {
		return err
	}
	return nil
}

func (s questionServiceImpl) GetAllQuestions() (dto.QuestionsResponse, error) {
	var questionsDto dto.QuestionsResponse

	questions, err := s.questionQuery.GetAllQuestionWithOwner()
	if err != nil {
		return questionsDto, err
	}

	for _, q := range questions {
		var questionDto dto.QuestionResponse
		questionDto.ID = q.ID
		questionDto.Question = q.Question
		questionDto.TotalVote = q.TotalVote
		questionDto.CreatedAt = q.CreatedAt
		questionDto.Owner.ID = q.User.ID
		questionDto.Owner.Name = q.User.Name
		questionDto.Owner.Username = q.User.Username
		questionsDto = append(questionsDto, questionDto)
	}

	return questionsDto, nil
}

func NewQuestionService(qr repository.QuestionRepository, qq query.QuestionQuery) *questionServiceImpl {
	return &questionServiceImpl{questionRepo: qr, questionQuery: qq}
}
