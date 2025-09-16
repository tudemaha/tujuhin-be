package dto

type QuestionRequestBody struct {
	Question string `json:"question" validate:"required"`
}
