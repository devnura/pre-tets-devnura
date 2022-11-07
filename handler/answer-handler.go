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
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AnswerHandler struct {
	answerService service.AnswerService
}

func NewAnswerHandler(service service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		answerService: service,
	}
}

// Answer godoc
// @Summary Answer
// @Description Answer
// @Tags answer
// @Accept  json
// @Produce  json
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /answer [get]
func (handler *AnswerHandler) AllAnswer(ctx echo.Context) (err error) {
	var data []entity.Answer = handler.answerService.AllAnswer()
	response := helper.BuildResponse(http.StatusOK, "OK!", data)

	return ctx.JSON(http.StatusOK, response)
}

// Answer godoc
// @Summary Answer By ID
// @Description Answer By ID
// @Tags answer
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Answer"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /answer/{id} [get]
func (c *AnswerHandler) FindAnswerById(ctx echo.Context) (err error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", "", []helper.EmptyObj{})

		return ctx.JSON(http.StatusBadRequest, response)
	}

	var answer entity.Answer = c.answerService.FindAnswerById(id)
	if (answer == entity.Answer{}) {
		res := helper.BuildErrorResponse(http.StatusNotFound, "Data not found", "No data with given id", helper.EmptyObj{})
		return ctx.JSON(http.StatusNotFound, res)
	}

	res := helper.BuildResponse(http.StatusOK, "OK!", answer)
	return ctx.JSON(http.StatusOK, res)
}

// Answer godoc
// @Summary Answer Insert
// @Description Answer Insert
// @Tags answer
// @Accept  json
// @Produce  json
// @param register body dto.AnswerRequestDTO true "request body insert question"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /answer [post]
func (c *AnswerHandler) InsertAnswer(ctx echo.Context) (err error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	validate := validator.New()

	var answerRequestDTO dto.AnswerRequestDTO
	defer cancel()

	errDTO := ctx.Bind(&answerRequestDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)

	}
	//use the validator library to validate required fields
	if err := validate.Struct(&answerRequestDTO); err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	var answerCreateDTO dto.AnswerCreateDTO
	answerCreateDTO.Answer = answerRequestDTO.Answer
	answerCreateDTO.QuestionID = answerRequestDTO.QuestionID
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		answerCreateDTO.UserID = convertedUserID
	} else {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result := c.answerService.InsertAnswer(answerCreateDTO)
	response := helper.BuildResponse(http.StatusCreated, "Created", result)
	return ctx.JSON(http.StatusCreated, response)
}

// Answwer godoc
// @Summary Update Answwer By ID
// @Description Update Answwer By ID
// @Tags answer
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Answer"
// @param register body dto.QuestionRequestDTO true "request body insert answer"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /answer/{id} [put]
func (c *AnswerHandler) UpdateAnswer(ctx echo.Context) (err error) {
	var answerUpdateDTO dto.AnswerUpdateDTO
	errDTO := ctx.Bind(&answerUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)

	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	questionID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err == nil {
		answerUpdateDTO.ID = questionID
	}

	if c.answerService.IsAllowedToEditAnswer(userID, answerUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			answerUpdateDTO.UserID = id
		}
		result := c.answerService.UpdateAnswer(answerUpdateDTO)
		response := helper.BuildResponse(http.StatusOK, "OK!", result)

		return ctx.JSON(http.StatusOK, response)
	} else {
		res := helper.BuildErrorResponse(http.StatusUnauthorized, "You dont have permission", "You are not the owner", helper.EmptyObj{})
		return ctx.JSON(http.StatusForbidden, res)
	}

}

// Answer godoc
// @Summary Delete Answer By ID
// @Description Delete Answer By ID
// @Tags answer
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Answer"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /answer/{id} [delete]
func (c *AnswerHandler) DeleteAnswer(ctx echo.Context) (err error) {
	var answer entity.Answer
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		res := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to get param ID", "Please insert param ID", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	if c.answerService.IsAllowedToEditAnswer(userID, id) {
		c.answerService.DeleteAnswer(answer)
		response := helper.BuildResponse(http.StatusOK, "OK!", helper.EmptyObj{})
		return ctx.JSON(http.StatusOK, response)
	} else {
		res := helper.BuildErrorResponse(http.StatusUnauthorized, "You dont have permission", "You are not the owner", helper.EmptyObj{})
		return ctx.JSON(http.StatusForbidden, res)
	}

}
