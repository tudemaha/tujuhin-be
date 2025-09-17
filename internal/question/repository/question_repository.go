package repository

import "github.com/tudemaha/tujuhin-be/internal/question/model"

type QuestionRepository interface {
	CreateQuestion(qm model.QuestionModel) error
	GetQuestions() ([]model.QuestionModel, error)
}
