package models

type ProductInput struct {
	Id            int64  `gorm:"primaryKey" json:"id"`
	Image         string `gorm:"type:longtext;not null" json:"image"`
	ProductName   string `gorm:"type:varchar(255);not null" json:"product_name"`
	Description   string `gorm:"type:text;not null" json:"description"`
	Price         int64  `gorm:"type:int;not null" json:"price"`
	SubcategoryID int64  `json:"sub_category_id"`
	ColorID       int64  `json:"color_id"`
}

type Product struct {
	Id            int64       `gorm:"primaryKey" json:"id"`
	Image         string      `gorm:"type:longtext;not null" json:"image"`
	ProductName   string      `gorm:"type:varchar(255);not null" json:"product_name"`
	Description   string      `gorm:"type:text;not null" json:"description"`
	Price         int64       `gorm:"type:int;not null" json:"price"`
	SubcategoryID int64       `json:"sub_category_id"`
	Subcategory   Subcategory `gorm:"foreignKey:SubcategoryID"`
}

type Product_Color struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	ProductID int64   `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	ColorID   int64   `json:"color_id"`
	Color     Color   `gorm:"foreignKey:ColorID"`
}