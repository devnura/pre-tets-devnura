package service

import (
	"log"

	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPass string, plainPassword []byte) bool {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
