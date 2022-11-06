package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}

type LoginResponseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
