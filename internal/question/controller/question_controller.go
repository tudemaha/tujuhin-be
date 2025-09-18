package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
	"github.com/tudemaha/tujuhin-be/internal/question/service"
	"github.com/tudemaha/tujuhin-be/pkg/dto/response"
	"github.com/tudemaha/tujuhin-be/pkg/server/middleware"
	"github.com/tudemaha/tujuhin-be/pkg/utils"
)

type QuestionController struct {
	questionGroup   *gin.RouterGroup
	questionService service.QuestionService
	authMiddleware  middleware.AuthMiddleware
}

func (qc *QuestionController) handleNewQuestion() gin.HandlerFunc {
	return func(c *gin.Context) {
		var question dto.QuestionRequestBody
		var baseResponse response.BaseResponse

		if err := c.Bind(&question); err != nil {
			baseResponse.DefaultInternalError()
			resErr := response.NewErrorResponseValue("bind_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(resErr)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		if arrError, err := utils.RequestBodyValidator(question); err {
			baseResponse.DefaultBadRequest()
			baseResponse.Errors = arrError
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		userID := c.MustGet("userID").(string)
		if err := qc.questionService.CreateQuestion(question, userID); err != nil {
			baseResponse.DefaultBadRequest()
			resErr := response.NewErrorResponseValue("insert_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(resErr)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		baseResponse.DefaultCreated()
		c.JSON(baseResponse.Code, baseResponse)
	}
}

func (qc *QuestionController) handleGetQuestions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var baseResponse response.BaseResponse
		userID := c.MustGet("userID").(string)

		questions, err := qc.questionService.GetAllQuestions(userID)
		if err != nil {
			baseResponse.DefaultInternalError()
			errRes := response.NewErrorResponseValue("get_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(errRes)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		baseResponse.DefaultOK()
		baseResponse.Data = gin.H{"questions": questions}
		c.JSON(baseResponse.Code, baseResponse)
	}
}

func (qc *QuestionController) handleVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var baseResponse response.BaseResponse
		var body dto.UpVoteBody
		userID := c.MustGet("userID").(string)

		if err := c.Bind(&body); err != nil {
			baseResponse.DefaultInternalError()
			resErr := response.NewErrorResponseValue("bind_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(resErr)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		if arrError, err := utils.RequestBodyValidator(body); err {
			baseResponse.DefaultBadRequest()
			baseResponse.Errors = arrError
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		if err := qc.questionService.Vote(body.QuestionID, userID, body.Vote); err != nil {
			baseResponse.DefaultBadRequest()
			resErr := response.NewErrorResponseValue("update_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(resErr)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		baseResponse.DefaultOK()
		baseResponse.Message = fmt.Sprintf("success vote %s question", body.Vote)
		c.JSON(baseResponse.Code, baseResponse)
	}
}

func (qc *QuestionController) InitializeController() {
	qc.questionGroup.POST("", qc.authMiddleware.Auth(), qc.handleNewQuestion())
	qc.questionGroup.GET("", qc.authMiddleware.Auth(), qc.handleGetQuestions())
	qc.questionGroup.PATCH("/vote", qc.authMiddleware.Auth(), qc.handleVote())
}

func NewQuestionController(rg *gin.RouterGroup, qs service.QuestionService, am middleware.AuthMiddleware) *QuestionController {
	return &QuestionController{questionGroup: rg, questionService: qs, authMiddleware: am}
}
