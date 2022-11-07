package repository

import (
	"github.com/devnura/pre-tets-devnura/entity"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	FindById(answerID uint64) entity.Answer
	AllAnswer() []entity.Answer
	InsertAnswer(b entity.Answer) entity.Answer
	UpdateAnswer(b entity.Answer) entity.Answer
	DeleteAnswer(b entity.Answer)
}

type answerConnection struct {
	connection *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerConnection{
		connection: db,
	}
}

func (db *answerConnection) FindById(answerID uint64) entity.Answer {
	var answer entity.Answer
	db.connection.Find(&answer, answerID)
	return answer
}

func (db *answerConnection) AllAnswer() []entity.Answer {
	var answer []entity.Answer
	db.connection.Find(&answer)
	return answer
}

func (db *answerConnection) InsertAnswer(b entity.Answer) entity.Answer {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *answerConnection) UpdateAnswer(b entity.Answer) entity.Answer {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *answerConnection) DeleteAnswer(b entity.Answer) {
	db.connection.Delete(&b)
}
