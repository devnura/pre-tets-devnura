package entity

import "time"

type Answer struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"i_id"`
	Question   string    `gorm:"type:varchar(255)" json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	User       User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
