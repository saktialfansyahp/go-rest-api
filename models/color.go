package models

type Color struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Color string `gorm:"type:varchar(255);not null" json:"color"`
}