package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	authController "github.com/tudemaha/tujuhin-be/internal/auth/controller"
	userRepo "github.com/tudemaha/tujuhin-be/internal/auth/repository"
	authService "github.com/tudemaha/tujuhin-be/internal/auth/service"
	questionController "github.com/tudemaha/tujuhin-be/internal/question/controller"
	questionQuery "github.com/tudemaha/tujuhin-be/internal/question/query"
	questionRepo "github.com/tudemaha/tujuhin-be/internal/question/repository"
	questionService "github.com/tudemaha/tujuhin-be/internal/question/service"
	"github.com/tudemaha/tujuhin-be/pkg/hasher"
	"github.com/tudemaha/tujuhin-be/pkg/jwt"
	"github.com/tudemaha/tujuhin-be/pkg/server/middleware"
)

func InitializeControllers(r *gin.Engine, db *sqlx.DB) {

	authRoutes := r.Group("/auth")
	hasher := hasher.BcryptHasher{}
	jwtPkg := jwt.JWT{}
	authRepoImpl := userRepo.NewAuthRepository(db)
	authServiceImpl := authService.NewAuthService(authRepoImpl, hasher, jwtPkg)
	authMiddleware := middleware.NewAuthMiddleware(authServiceImpl)
	authControllerImpl := authController.NewAuthController(authRoutes, authServiceImpl)
	authControllerImpl.InitializeController()

	questionRoutes := r.Group("/questions")
	questionRepoImpl := questionRepo.NewQuestionRepository(db)
	questionQeuryImpl := questionQuery.NewQueryQuery(db)
	questionServiceImpl := questionService.NewQuestionService(questionRepoImpl, questionQeuryImpl)
	questionControllerImpl := questionController.NewQuestionController(questionRoutes, questionServiceImpl, authMiddleware)
	questionControllerImpl.InitializeController()
}
