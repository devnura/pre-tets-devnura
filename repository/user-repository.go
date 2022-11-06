package repository

import (
	"log"

	"github.com/devnura/pre-tets-devnura/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	VerifyCredential(email string, password string) interface{}
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
