package repository

import "github.com/tudemaha/tujuhin-be/internal/question/model"

type QuestionRepository interface {
	CreateQuestion(qm model.Question) error
}
