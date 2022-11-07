package dto

type AnswerUpdateDTO struct {
	ID         uint64 `json:"id" form:"id"`
	Answer     string `json:"answer" validate:"required"`
	UserID     uint64 `json:"userId"`
	QuestionID uint64 `json:"questionId" validate:"required"`
}

type AnswerCreateDTO struct {
	Answer     string `json:"answer" validate:"required"`
	UserID     uint64 `json:"userId"`
	QuestionID uint64 `json:"questionId" validate:"required"`
}

type AnswerRequestDTO struct {
	Answer     string `json:"answer" validate:"required"`
	QuestionID uint64 `json:"questionId" validate:"required"`
}
