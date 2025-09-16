package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/tujuhin-be/internal/auth/service"
	"github.com/tudemaha/tujuhin-be/pkg/dto/response"
)

type authMiddleware struct {
	authService service.AuthService
}

func (am authMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawAuthHeader := strings.Split(c.GetHeader("Authorization"), " ")
		var baseResponse response.BaseResponse

		if len(rawAuthHeader) < 2 {
			baseResponse.DefaultBadRequest()
			errRes := response.NewErrorResponseValue("auth_error", "Authorization header required")
			baseResponse.Errors = response.NewArrErrorResponse(errRes)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		userID, err := am.authService.ValidateToken(rawAuthHeader[1])
		if err != nil {
			baseResponse.DefaultUnauthorized()
			errRes := response.NewErrorResponseValue("auth_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(errRes)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

func NewAuthMiddleware(as service.AuthService) *authMiddleware {
	return &authMiddleware{authService: as}
}
