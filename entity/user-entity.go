package entity

import "time"

type User struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"passwrod_hash"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	Token      string    `gorm:"-" json:"token,omitempty"`
}

type Users []User
