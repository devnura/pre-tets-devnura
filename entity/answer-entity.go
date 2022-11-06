package entity

type Answer struct {
	ID         uint64   `gorm:"primary_key:auto_increment" json:"i_id"`
	Answer     string   `gorm:"type:varchar(255)" json:"title"`
	QuestionID uint64   `json:"question_id"`
	Question   Question `gorm:"foreignkey:QuestionID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"question"`
}
