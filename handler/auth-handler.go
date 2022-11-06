package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/devnura/pre-tets-devnura/dto"
	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/helper"
	"github.com/devnura/pre-tets-devnura/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// type AuthHandler interface {
// 	Login(ctx echo.Context)
// }

type AuthHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthHandler(service service.AuthService, jwtService service.JWTService) *AuthHandler {
	return &AuthHandler{
		authService: service,
		jwtService:  jwtService,
	}
}

func (c *AuthHandler) Login(ctx echo.Context) (err error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	validate := validator.New()

	var loginDTO dto.LoginDTO
	defer cancel()

	errDTO := ctx.Bind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), []helper.EmptyObj{})

		return ctx.JSON(http.StatusBadRequest, response)
	}

	//use the validator library to validate required fields
	if err := validate.Struct(&loginDTO); err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		// generatedToken := "abcg"
		v.Token = generatedToken

		response := helper.BuildResponse(http.StatusOK, "OK!", dto.LoginResponseDTO{Name: v.Name, Email: v.Email, Token: v.Token})

		return ctx.JSON(http.StatusOK, response)
	}

	response := helper.BuildErrorResponse(http.StatusUnauthorized, "Unauthorized", "Invalid email or password", []helper.EmptyObj{})

	return ctx.JSON(http.StatusUnauthorized, response)
}
