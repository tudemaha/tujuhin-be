package service

import (
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
)

type QuestionService interface {
	CreateQuestion(qd dto.QuestionRequestBody, owner string) error
	GetAllQuestions(userID string) (dto.QuestionsResponse, error)
	Vote(questionID, userID, newVote string) error
}
