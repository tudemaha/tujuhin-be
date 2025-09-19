package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
	"github.com/tudemaha/tujuhin-be/internal/question/model"
	"github.com/tudemaha/tujuhin-be/internal/question/query"
	"github.com/tudemaha/tujuhin-be/internal/question/repository"
)

type questionServiceImpl struct {
	questionRepo  repository.QuestionRepository
	questionQuery query.QuestionQuery
}

func (s questionServiceImpl) CreateQuestion(qd dto.QuestionRequestBody, owner string) error {
	userID := uuid.MustParse(owner)
	question := model.QuestionModel{
		UserID:   userID,
		Question: qd.Question,
	}

	if err := s.questionRepo.CreateQuestion(question); err != nil {
		return err
	}
	return nil
}

func (s questionServiceImpl) GetAllQuestions(userID string) (dto.QuestionsResponse, error) {
	var questionsDto dto.QuestionsResponse

	questions, err := s.questionQuery.GetAllQuestionWithOwner(userID)
	if err != nil {
		return questionsDto, err
	}

	for _, q := range questions {
		var questionDto dto.QuestionResponse
		questionDto.ID = q.ID
		questionDto.Question = q.Question
		questionDto.TotalVote = q.TotalVote
		questionDto.CreatedAt = q.CreatedAt
		questionDto.Owner.ID = q.User.ID
		questionDto.Owner.Name = q.User.Name
		questionDto.Owner.Username = q.User.Username
		if q.Vote.VoteState != nil {
			questionDto.VoteState = *q.Vote.VoteState
		}
		questionsDto = append(questionsDto, questionDto)
	}

	return questionsDto, nil
}

func (s questionServiceImpl) Vote(questionID, userID, newVote string) error {
	vote, err := s.questionRepo.GetVoteByQuestionUser(questionID, userID)
	if err != nil {
		return err
	}

	var voteStr string
	if vote.VoteState != nil {
		voteStr = *vote.VoteState
	}

	if voteStr == "" {
		vote.QuestionID = uuid.MustParse(questionID)
		vote.UserID = uuid.MustParse(userID)
		vote.VoteState = &newVote
		if err := s.questionRepo.CreateNewVote(vote); err != nil {
			return err
		}
	} else if voteStr != newVote {
		vote.VoteState = &newVote
		if err := s.questionRepo.UpdateVoteByID(vote); err != nil {
			return err
		}
	} else {
		return errors.New("already vote for current question")
	}

	if err := s.updateVote(questionID); err != nil {
		return err
	}

	return nil
}

func (s questionServiceImpl) DeleteVote(questionID, userID string) error {
	vote, err := s.questionRepo.GetVoteByQuestionUser(questionID, userID)
	if err != nil {
		return err
	}

	if vote.VoteState == nil {
		return errors.New("vote not found for current question")
	}

	if err := s.questionRepo.DeleteVoteByID(vote.ID.String()); err != nil {
		return err
	}

	if err := s.updateVote(questionID); err != nil {
		return err
	}

	return nil
}

func (s questionServiceImpl) updateVote(questionID string) error {
	totalVote, err := s.questionRepo.GetTotalVote(questionID)
	if err != nil {
		return err
	}

	if err := s.questionRepo.UpdateTotalVote(totalVote, questionID); err != nil {
		return err
	}

	return nil
}

func NewQuestionService(qr repository.QuestionRepository, qq query.QuestionQuery) *questionServiceImpl {
	return &questionServiceImpl{questionRepo: qr, questionQuery: qq}
}
