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
	DeleteQuestion(b entity.Question)
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
	db.connection.Preload("Answer").Preload("Answe.Question").Find(&question, questionID)
	return question
}

func (db *questionConnection) AllQuestion() []entity.Question {
	var entity []entity.Question
	db.connection.Raw("SELECT questions.id, question, users.* FROM questions LEFT JOIN users ON users.id = questions.user_id ").Scan(&entity)
	return entity
}

func (db *questionConnection) InsertQuestion(b entity.Question) entity.Question {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *questionConnection) UpdateQuestion(b entity.Question) entity.Question {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *questionConnection) DeleteQuestion(b entity.Question) {
	db.connection.Delete(&b)
}
