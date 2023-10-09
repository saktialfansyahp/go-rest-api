package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
    // db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi"))
    // dsn := "host=localhost user=postgres password=alfan12345 dbname=asaba port=5432"
    dsn := "postgres://default:xBTaUne1kOj2@ep-shrill-base-66245357.ap-southeast-1.postgres.vercel-storage.com:5432/verceldb"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if(err != nil){
        panic("failed to connect database")
    }

	db.AutoMigrate(&User{}, &Role{}, &Product{}, &Category{}, &Subcategory{}, &Color{}, &Product_Color{})
    
    DB = db
}