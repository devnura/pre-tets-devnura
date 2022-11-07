package repository

import (
	"github.com/devnura/pre-tets-devnura/entity"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	FindById(questionID uint64) entity.Question
	AllQuestion() []entity.Question
	InsertQuestion(b entity.Question) entity.Question
	UpdateQuestion(b entity.Question) entity.Question
	DeleteQuestion(b entity.Question, questionID uint64)
	FindAnswer(questionID uint64) []entity.Answer
}

type questionConnection struct {
	connection *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionConnection{
		connection: db,
	}
}

func (db *questionConnection) FindById(questionID uint64) entity.Question {
	var question entity.Question
	db.connection.Find(&question, questionID)
	return question
}

func (db *questionConnection) FindAnswer(questionID uint64) []entity.Answer {
	var answer []entity.Answer
	db.connection.Find(&answer, questionID)
	return answer
}

func (db *questionConnection) AllQuestion() []entity.Question {
	var question []entity.Question
	db.connection.Find(&question)
	return question
}

func (db *questionConnection) InsertQuestion(b entity.Question) entity.Question {
	db.connection.Save(&b)
	db.connection.Preload("Answer").Find(&b)
	return b
}

func (db *questionConnection) UpdateQuestion(b entity.Question) entity.Question {
	db.connection.Save(&b)
	db.connection.Preload("Answer").Find(&b)
	return b
}

func (db *questionConnection) DeleteQuestion(b entity.Question, questionID uint64) {
	var answer entity.Answer
	db.connection.Where("question_id = ?", questionID).Delete(&answer)
	db.connection.Where("id = ?", questionID).Delete(&b)
}
