package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tudemaha/tujuhin-be/internal/auth/dto"
	"github.com/tudemaha/tujuhin-be/internal/auth/service"
	"github.com/tudemaha/tujuhin-be/pkg/dto/response"
	"github.com/tudemaha/tujuhin-be/pkg/utils"
)

type AuthController struct {
	authGroup   *gin.RouterGroup
	authService service.AuthService
}

func (ac *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.UserRegister
		var baseResponse response.BaseResponse

		if err := c.Bind(&user); err != nil {
			baseResponse.DefaultInternalError()
			resErr := response.NewErrorResponseValue("bind_error", err.Error())
			baseResponse.Errors = response.NewArrErrorResponse(resErr)
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		if arrError, err := utils.RequestBodyValidator(user); err {
			baseResponse.DefaultBadRequest()
			baseResponse.Errors = arrError
			c.AbortWithStatusJSON(baseResponse.Code, baseResponse)
			return
		}

		if err := ac.authService.Register(user); err != nil {
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

func (ac *AuthController) InitializeController() {
	ac.authGroup.POST("/register", ac.Register())
}

func ProvideAuthController(rg *gin.RouterGroup, as service.AuthService) *AuthController {
	return &AuthController{authGroup: rg, authService: as}
}
