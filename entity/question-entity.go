package entity

type Question struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"id"`
	Question string `gorm:"type:varchar(255)" json:"question"`
	UserID   int    `gorm:"type:int(11);not null" json:"user_id"`
	User     User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
