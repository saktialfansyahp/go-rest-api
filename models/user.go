package models

type AuthRequest struct {
	// gorm.Model
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password string `gorm:"type:varchar(300)" json:"password"`
	RoleID   int64  `json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

type User struct {
	// gorm.Model
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password string `gorm:"type:varchar(300)" json:"password"`
	RoleID   int64  `json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	// gorm.Model
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
}