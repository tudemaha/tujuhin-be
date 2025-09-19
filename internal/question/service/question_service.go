package service

import (
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
)

type QuestionService interface {
	CreateQuestion(qd dto.QuestionRequestBody, owner string) error
	GetAllQuestions() (dto.QuestionsResponse, error)
	Vote(questionID, userID, newVote string) error
	DeleteVote(questionID, userID string) error
}
