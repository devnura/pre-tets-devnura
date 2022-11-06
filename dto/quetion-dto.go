package dto

type QuestionUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Question string `json:"question" validate:"required,email"`
	UserID   string `json:"userId" validate:"required"`
}

type QuestionCreateDTO struct {
	Question string `json:"question" validate:"required"`
	UserID   uint64 `json:"userId"`
}
type QuestionRequestDTO struct {
	Question string `json:"question" validate:"required"`
}
