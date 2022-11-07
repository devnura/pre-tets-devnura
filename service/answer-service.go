package service

import (
	"fmt"
	"log"

	"github.com/devnura/pre-tets-devnura/dto"
	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/repository"
	"github.com/mashingan/smapping"
)

type AnswerService interface {
	InsertAnswer(b dto.AnswerCreateDTO) entity.Answer
	UpdateAnswer(b dto.AnswerUpdateDTO) entity.Answer
	DeleteAnswer(b entity.Answer)
	AllAnswer() []entity.Answer
	FindAnswerById(AnswerID uint64) entity.Answer
	IsAllowedToEditAnswer(userID string, answerID uint64) bool
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func NewAnswerService(answerRepo repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepository: answerRepo,
	}
}

func (service *answerService) InsertAnswer(b dto.AnswerCreateDTO) entity.Answer {
	data := entity.Answer{}
	err := smapping.FillStruct(&data, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.answerRepository.InsertAnswer(data)
	return res
}

func (service *answerService) UpdateAnswer(b dto.AnswerUpdateDTO) entity.Answer {
	data := entity.Answer{}
	err := smapping.FillStruct(&data, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.answerRepository.UpdateAnswer(data)
	return res
}

func (service *answerService) DeleteAnswer(b entity.Answer) {
	service.answerRepository.DeleteAnswer(b)
}

func (service *answerService) AllAnswer() []entity.Answer {
	return service.answerRepository.AllAnswer()
}

func (service *answerService) FindAnswerById(answerID uint64) entity.Answer {
	return service.answerRepository.FindById(answerID)
}

func (service *answerService) IsAllowedToEditAnswer(UserID string, answerID uint64) bool {
	b := service.answerRepository.FindById(answerID)
	id := fmt.Sprintf("%v", b.UserID)
	return UserID == id
}
