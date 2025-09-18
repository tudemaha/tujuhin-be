package repository

import (
	"github.com/tudemaha/tujuhin-be/internal/question/model"
)

type QuestionRepository interface {
	CreateQuestion(qm model.QuestionModel) error
	UpdateTotalVote(totalVote int, questionID string) error
	CreateNewVote(vm model.QuestionVote) error
	UpdateVoteByID(vm model.QuestionVote) error
	GetVoteByQuestionUser(questionID, userID string) (model.QuestionVote, error)
	GetTotalVote(questionID string) (int, error)
}
