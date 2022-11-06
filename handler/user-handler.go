package handler

import (
	"net/http"

	"github.com/devnura/pre-tets-devnura/dto"
	"github.com/devnura/pre-tets-devnura/helper"
	"github.com/devnura/pre-tets-devnura/service"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// Login godoc
// @Summary Provole
// @Description Provole
// @Tags auth
// @Accept  json
// @Produce  json
// @param register body LoginDTO true "request body login"
// @Success 200 {object} LoginResponseDOC
// @Failure 400 {object} res.ErrorResponse
// @Failure 404 {object} res.ErrorResponse
// @Failure 500 {object} res.ErrorResponse
// @Router /auth/login [post]

func (handler *UserHandler) Profile(ctx echo.Context) (err error) {
	// c.Request().Header
	var userDTO dto.UserDTO
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	userProfile := handler.userService.Profile(userId)
	errDTO := copier.Copy(&userDTO, userProfile)
	if err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), []helper.EmptyObj{})

		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := helper.BuildResponse(http.StatusOK, "OK!", userProfile)
	return ctx.JSON(http.StatusOK, response)
}
