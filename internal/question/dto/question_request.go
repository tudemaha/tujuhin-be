package dto

type QuestionRequestBody struct {
	Question string `json:"question" validate:"required"`
}

type UpVoteBody struct {
	QuestionID string `json:"question_id" validate:"required"`
	Vote       string `json:"vote" validate:"required,oneof=up down"`
}
