package service

import (
	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/repository"
)

type UserService interface {
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
