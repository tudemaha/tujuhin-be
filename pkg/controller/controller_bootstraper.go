package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tudemaha/tujuhin-be/internal/question/controller"
	"github.com/tudemaha/tujuhin-be/internal/question/repository"
	"github.com/tudemaha/tujuhin-be/internal/question/service"
)

func InitializeControllers(r *gin.Engine, db *sqlx.DB) {
	questionRoutes := r.Group("/questions")
	questionRepo := repository.ProvideQuestionRepository(db)
	questionService := service.ProvideQuestionService(questionRepo)
	questionController := controller.ProvideQuestionController(questionRoutes, questionService)
	questionController.InitializeController()
}
