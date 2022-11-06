package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/devnura/pre-tets-devnura/dto"
	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/helper"
	"github.com/devnura/pre-tets-devnura/service"
	"github.com/labstack/echo/v4"
)

// type AuthHandler interface {
// 	Login(ctx echo.Context)
// }

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

func (c *authHandler) Login(ctx echo.Context) (err error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var loginDTO dto.LoginDTO
	defer cancel()
	errDTO := ctx.Bind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		// generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		generatedToken := "abcg"
		v.Token = generatedToken

		response := helper.BuildResponse(true, "OK!", dto.LoginResponseDTO{Name: v.Name, Email: v.Email, Token: v.Token})
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Failed to login", "Invalid credentials", authResult)
	ctx.JSON(http.StatusUnauthorized, response)
	return
}
