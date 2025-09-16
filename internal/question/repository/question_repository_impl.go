package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
)

type questionRepositoryImpl struct {
	DB *sqlx.DB
}

func (q questionRepositoryImpl) CreateQuestion(qm model.Question) error {
	stmt, err := q.DB.Preparex("INSERT INTO questions (user_id, question) VALUES (?, ?);")
	if err != nil {
		return err
	}

	stmt.Exec(qm.UserID, qm.Question)

	return nil
}

func ProvideQuestionRepository(DB *sqlx.DB) *questionRepositoryImpl {
	return &questionRepositoryImpl{DB: DB}
}
