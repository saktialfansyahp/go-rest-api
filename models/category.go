package models

type Category struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Category    string `gorm:"type:varchar(255)" json:"category"`
	Description string `gorm:"type:text" json:"description"`
}