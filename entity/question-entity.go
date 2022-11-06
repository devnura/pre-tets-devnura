package entity

import "time"

type Question struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"i_id"`
	UserID     uint64    `json:"i_user_id"`
	Question   string    `gorm:"type:varchar(255)" json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	User       User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Books      *[]Answer `json:"answer,omitempty"`
}
