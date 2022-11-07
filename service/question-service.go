package service

import (
	"fmt"
	"log"

	"github.com/devnura/pre-tets-devnura/dto"
	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/devnura/pre-tets-devnura/repository"
	"github.com/mashingan/smapping"
)

type QuestionService interface {
	Insert(b dto.QuestionCreateDTO) entity.Question
	Update(b dto.QuestionUpdateDTO) entity.Question
	Delete(b entity.Question, questionID uint64)
	All() []entity.Question
	FindById(QuestionID uint64) entity.Question
	IsAllowedToEdit(userID string, questionID uint64) bool
}

type questionService struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionService(questionRepo repository.QuestionRepository) QuestionService {
	return &questionService{
		questionRepository: questionRepo,
	}
}

func (service *questionService) Insert(b dto.QuestionCreateDTO) entity.Question {
	data := entity.Question{}
	err := smapping.FillStruct(&data, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.questionRepository.InsertQuestion(data)
	return res
}

func (service *questionService) Update(b dto.QuestionUpdateDTO) entity.Question {
	data := entity.Question{}
	err := smapping.FillStruct(&data, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.questionRepository.UpdateQuestion(data)
	return res
}

func (service *questionService) Delete(b entity.Question, questionID uint64) {
	service.questionRepository.DeleteQuestion(b, questionID)
}

func (service *questionService) All() []entity.Question {
	return service.questionRepository.AllQuestion()
}

func (service *questionService) FindById(questionID uint64) entity.Question {
	return service.questionRepository.FindById(questionID)
}

func (service *questionService) IsAllowedToEdit(UserID string, questionID uint64) bool {
	b := service.questionRepository.FindById(questionID)
	id := fmt.Sprintf("%v", b.UserID)
	log.Printf("%s == %s", UserID, id)
	return UserID == id
}
