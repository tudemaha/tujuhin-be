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
	"github.com/tudemaha/tujuhin-be/pkg/hasher"
	"github.com/tudemaha/tujuhin-be/pkg/jwt"
)

func InitializeControllers(r *gin.Engine, db *sqlx.DB) {

	authRoutes := r.Group("/auth")
	hasher := hasher.BcryptHasher{}
	jwtPkg := jwt.JWT{}
	authRepoImpl := userRepo.NewAuthRepository(db)
	authServiceImpl := authService.NewAuthService(authRepoImpl, hasher, jwtPkg)
	authControllerImpl := authController.NewAuthController(authRoutes, authServiceImpl)
	authControllerImpl.InitializeController()

	questionRoutes := r.Group("/questions")
	questionRepoImpl := questionRepo.ProvideQuestionRepository(db)
	questionServiceImpl := questionService.ProvideQuestionService(questionRepoImpl)
	questionControllerImpl := questionController.ProvideQuestionController(questionRoutes, questionServiceImpl)
	questionControllerImpl.InitializeController()
}
