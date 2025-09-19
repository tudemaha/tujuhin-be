package query

import (
	"github.com/tudemaha/tujuhin-be/internal/question/model"
)

type QuestionQuery interface {
	GetAllQuestionWithOwner() (model.QuestionsWithOwner, error)
}
