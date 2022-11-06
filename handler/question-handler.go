package handler

import (
	"net/http"

	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/helper"
	"github.com/devnura/pre-tets-devnura/service"
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
