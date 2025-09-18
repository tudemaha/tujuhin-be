package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
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

func (q questionRepositoryImpl) UpdateTotalVote(totalVote int, questionID string) error {
	QID := uuid.MustParse(questionID)
	stmt, err := q.DB.Preparex(`UPDATE questions SET total_vote = $1 WHERE id = $2`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(totalVote, QID); err != nil {
		return err
	}

	return nil
}

func (q questionRepositoryImpl) GetVoteByQuestionUser(questionID, userID string) (model.QuestionVote, error) {
	QID := uuid.MustParse(questionID)
	UID := uuid.MustParse(userID)
	var vote model.QuestionVote

	stmt := `SELECT id, question_id, user_id, vote_state FROM question_votes WHERE question_id = $1 AND user_id = $2 LIMIT 1`
	if err := q.DB.Get(&vote, stmt, QID, UID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return vote, err
		}
	}

	return vote, nil
}

func (q questionRepositoryImpl) CreateNewVote(vm model.QuestionVote) error {
	stmt, err := q.DB.Preparex(`INSERT INTO question_votes (question_id, user_id, vote_state) VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(vm.QuestionID, vm.UserID, vm.VoteState); err != nil {
		return err
	}

	return nil
}

func (q questionRepositoryImpl) UpdateVoteByID(vm model.QuestionVote) error {
	stmt, err := q.DB.Preparex(`UPDATE question_votes SET vote_state = $1 WHERE id = $2`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(vm.VoteState, vm.ID); err != nil {
		return err
	}

	return nil
}

func (q questionRepositoryImpl) GetTotalVote(questionID string) (int, error) {
	QID := uuid.MustParse(questionID)
	stmt := `SELECT SUM(CASE WHEN vote_state = 'up' THEN 1 ELSE -1 END) 
		FROM question_votes WHERE question_id = $1`

	var totalVote int
	if err := q.DB.Get(&totalVote, stmt, QID); err != nil {
		return totalVote, err
	}

	return totalVote, nil
}

func NewQuestionRepository(DB *sqlx.DB) *questionRepositoryImpl {
	return &questionRepositoryImpl{DB: DB}
}
