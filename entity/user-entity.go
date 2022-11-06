package entity

type User struct {
	ID       uint64      `gorm:"type:int(11);primary_key:auto_increment" json:"id"`
	Name     string      `gorm:"type:varchar(255)" json:"name"`
	Email    string      `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string      `gorm:"->;<-;not null" json:"-"`
	Token    string      `gorm:"-" json:"token,omitempty"`
	Question *[]Question `json:"question,omitempty"`
}
