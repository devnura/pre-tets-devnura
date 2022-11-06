package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
