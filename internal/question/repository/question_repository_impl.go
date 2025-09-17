package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
)

type questionRepositoryImpl struct {
	DB *sqlx.DB
}

func (q questionRepositoryImpl) CreateQuestion(qm model.QuestionModel) error {
	stmt, err := q.DB.Preparex(`INSERT INTO questions (user_id, question) VALUES ($1, $2)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(qm.UserID, qm.Question); err != nil {
		return err
	}

	return nil
}

func (q questionRepositoryImpl) GetQuestions() ([]model.QuestionModel, error) {
	var questions []model.QuestionModel

	stmt := `SELECT id, user_id, questions, total_vote, created_at 
		FROM questions ORDER BY created_at DESC`

	if err := q.DB.Get(&questions, stmt); err != nil {
		return questions, err
	}

	return questions, nil
}

func NewQuestionRepository(DB *sqlx.DB) *questionRepositoryImpl {
	return &questionRepositoryImpl{DB: DB}
}
