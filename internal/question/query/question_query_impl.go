package query

import (
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
)

type questionQueryImpl struct {
	db *sqlx.DB
}

func (qq questionQueryImpl) GetAllQuestionWithOwner() (model.QuestionsWithOwner, error) {
	var questionsWithOwner model.QuestionsWithOwner
	stmt := `
		SELECT 
		q.id AS "id", q.question AS "question", q.total_vote AS "total_vote", q.created_at as "created_at", 
		u.id AS "user.id", u.name AS "user.name", u.username AS "user.username", 
		v.id AS "vote.id", v.vote_state AS "vote.vote_state" 
		FROM questions q 
		INNER JOIN users u ON u.id = q.user_id 
		LEFT JOIN question_votes v ON v.question_id = q.id AND v.user_id = u.id 
		ORDER BY q.created_at DESC 
	`

	err := qq.db.Select(&questionsWithOwner, stmt)
	if err != nil {
		return questionsWithOwner, err
	}

	return questionsWithOwner, nil
}

func NewQueryQuery(db *sqlx.DB) *questionQueryImpl {
	return &questionQueryImpl{db: db}
}
