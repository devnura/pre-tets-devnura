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

type QuestionHandler struct {
	questionService service.QuestionService
}

func NewQuestionHandler(service service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: service,
	}
}

// Question godoc
// @Summary Question
// @Description Question
// @Tags question
// @Accept  json
// @Produce  json
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /question [get]
func (handler *QuestionHandler) All(ctx echo.Context) (err error) {
	var data []entity.Question = handler.questionService.All()
	response := helper.BuildResponse(http.StatusOK, "OK!", data)

	return ctx.JSON(http.StatusOK, response)
}

// Question godoc
// @Summary Question By ID
// @Description Question By ID
// @Tags question
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Question"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /question/{id} [get]
func (c *QuestionHandler) FindById(ctx echo.Context) (err error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", "", []helper.EmptyObj{})

		return ctx.JSON(http.StatusBadRequest, response)
	}

	var question entity.Question = c.questionService.FindById(id)
	if (question == entity.Question{}) {
		res := helper.BuildErrorResponse(http.StatusNotFound, "Data not found", "No data with given id", helper.EmptyObj{})
		return ctx.JSON(http.StatusNotFound, res)
	}

	res := helper.BuildResponse(http.StatusOK, "OK!", question)
	return ctx.JSON(http.StatusOK, res)
}

// @Summary Get Answer ID Question
// @Description Get Answer By ID Question
// @Tags question
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Question"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /question/answer/{id} [get]
func (c *QuestionHandler) FindAnswer(ctx echo.Context) (err error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", "", []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	var question []entity.Answer = c.questionService.FindAnswer(id)
	response := helper.BuildResponse(http.StatusOK, "OK!", question)

	return ctx.JSON(http.StatusOK, response)
}

// Question godoc
// @Summary Question Insert
// @Description Question Insert
// @Tags question
// @Accept  json
// @Produce  json
// @param register body dto.QuestionRequestDTO true "request body insert question"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /question [post]
func (c *QuestionHandler) Insert(ctx echo.Context) (err error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	validate := validator.New()

	var questionRequestDTO dto.QuestionRequestDTO
	defer cancel()

	errDTO := ctx.Bind(&questionRequestDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)

	}
	//use the validator library to validate required fields
	if err := validate.Struct(&questionRequestDTO); err != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	var questionCreateDTO dto.QuestionCreateDTO
	questionCreateDTO.Question = questionRequestDTO.Question
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		questionCreateDTO.UserID = convertedUserID
	} else {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Bad Request", err.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result := c.questionService.Insert(questionCreateDTO)
	response := helper.BuildResponse(http.StatusCreated, "Created", result)
	return ctx.JSON(http.StatusCreated, response)
}

// Question godoc
// @Summary Update Question By ID
// @Description Update Question By ID
// @Tags question
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Question"
// @param register body dto.QuestionRequestDTO true "request body insert question"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /question/{id} [put]
func (c *QuestionHandler) Update(ctx echo.Context) (err error) {
	var questionUpdateDTO dto.QuestionUpdateDTO
	errDTO := ctx.Bind(&questionUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), []helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)

	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	questionID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err == nil {
		questionUpdateDTO.ID = questionID
	}

	if c.questionService.IsAllowedToEdit(userID, questionUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			questionUpdateDTO.UserID = id
		}
		result := c.questionService.Update(questionUpdateDTO)
		response := helper.BuildResponse(http.StatusOK, "OK!", result)

		return ctx.JSON(http.StatusOK, response)
	} else {
		res := helper.BuildErrorResponse(http.StatusUnauthorized, "You dont have permission", "You are not the owner", helper.EmptyObj{})
		return ctx.JSON(http.StatusForbidden, res)
	}

}

// Question godoc
// @Summary Delete Question By ID
// @Description Delete Question By ID
// @Tags question
// @Accept  json
// @Produce  json
// @Param        id    path      int     true  "Id Question"
// @Security  Bearer
// @Security   JWT
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /question/{id} [delete]
func (c *QuestionHandler) Delete(ctx echo.Context) (err error) {
	var question entity.Question

	questionID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err != nil {
		res := helper.BuildErrorResponse(http.StatusBadRequest, "Failed to get param ID", "Please insert param ID", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	if c.questionService.IsAllowedToEdit(userID, questionID) {
		c.questionService.Delete(question, questionID)
		response := helper.BuildResponse(http.StatusOK, "OK!", helper.EmptyObj{})
		return ctx.JSON(http.StatusOK, response)
	} else {
		res := helper.BuildErrorResponse(http.StatusUnauthorized, "You dont have permission", "You are not the owner", helper.EmptyObj{})
		return ctx.JSON(http.StatusForbidden, res)
	}
}
