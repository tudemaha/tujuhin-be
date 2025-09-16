package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tudemaha/tujuhin-be/internal/question/dto"
	"github.com/tudemaha/tujuhin-be/internal/question/service"
	"github.com/tudemaha/tujuhin-be/pkg/dto/response"
	"github.com/tudemaha/tujuhin-be/pkg/utils"
)

type QuestionController struct {
	questionGroup   *gin.RouterGroup
	questionService service.QuestionService
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

		if err := qc.questionService.CreateQuestion(question, uuid.MustParse("ffa13b22-8184-4a31-9bb9-12286a499153")); err != nil {
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

func (qc *QuestionController) InitializeController() {
	qc.questionGroup.POST("", qc.handleNewQuestion())
}

func ProvideQuestionController(rg *gin.RouterGroup, qs service.QuestionService) *QuestionController {
	return &QuestionController{questionGroup: rg, questionService: qs}
}
