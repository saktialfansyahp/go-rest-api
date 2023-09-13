package models

type Subcategory struct {
	Id          int64    `gorm:"primaryKey" json:"id"`
	Subcategory string   `gorm:"type:varchar(255)" json:"sub_category"`
	Description string   `gorm:"type:text" json:"description"`
	CategoryID  int64    `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}