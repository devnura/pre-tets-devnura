package entity

type Question struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Question string `gorm:"type:varchar(255)" json:"question"`
	UserID   uint64 `gorm:"not null" json:"-"`
}
