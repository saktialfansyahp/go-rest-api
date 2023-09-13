package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dbURL := "postgres://default:************@ep-shrill-base-66245357.ap-southeast-1.postgres.vercel-storage.com:5432/verceldb"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

	db.AutoMigrate(&User{}, &Role{}, &Product{}, &Category{}, &Subcategory{}, &Color{}, &Product_Color{})

    DB = db
}