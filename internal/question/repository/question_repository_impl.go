package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
)

type questionRepositoryImpl struct {
	DB *sqlx.DB
}

func (q questionRepositoryImpl) CreateQuestion(qm model.Question) error {
	stmt, err := q.DB.Preparex(`INSERT INTO questions (user_id, question) VALUES ($1, $2)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(qm.UserID, qm.Question); err != nil {
		return err
	}

	return nil
}

func ProvideQuestionRepository(DB *sqlx.DB) *questionRepositoryImpl {
	return &questionRepositoryImpl{DB: DB}
}
