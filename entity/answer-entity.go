package entity

type Answer struct {
	ID         int      `gorm:"primary_key:auto_increment" json:"id"`
	Answer     string   `gorm:"type:varchar(255)" json:"answer"`
	UserID     int      `gorm:"type:int(11);not null" json:"user_id"`
	QuestionID int      `gorm:"type:int(11);not null" json:"question_id"`
	User       User     `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Question   Question `gorm:"foreignkey:QuestionID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
