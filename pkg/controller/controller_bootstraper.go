package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	authController "github.com/tudemaha/tujuhin-be/internal/auth/controller"
	userRepo "github.com/tudemaha/tujuhin-be/internal/auth/repository"
	authService "github.com/tudemaha/tujuhin-be/internal/auth/service"
	questionController "github.com/tudemaha/tujuhin-be/internal/question/controller"
	questionRepo "github.com/tudemaha/tujuhin-be/internal/question/repository"
	questionService "github.com/tudemaha/tujuhin-be/internal/question/service"
)

func InitializeControllers(r *gin.Engine, db *sqlx.DB) {
	authRoutes := r.Group("/auth")
	authRepoImpl := userRepo.ProvideUserRepository(db)
	authServiceImpl := authService.ProvideUserRepository(authRepoImpl)
	authControllerImpl := authController.ProvideAuthController(authRoutes, authServiceImpl)
	authControllerImpl.InitializeController()

	questionRoutes := r.Group("/questions")
	questionRepoImpl := questionRepo.ProvideQuestionRepository(db)
	questionServiceImpl := questionService.ProvideQuestionService(questionRepoImpl)
	questionControllerImpl := questionController.ProvideQuestionController(questionRoutes, questionServiceImpl)
	questionControllerImpl.InitializeController()
}
