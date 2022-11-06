package repository

import (
	"log"

	"github.com/devnura/pre-tets-devnura/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	VerifyCredential(email string, password string) interface{}
	ProfileUser(id string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Printf("%v", res.Error)
		return nil
	}
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Question").Preload("Question.User").Find(&user, userID)
	return user
}
