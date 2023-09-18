package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	// dbURL := "postgres://default:xBTaUne1kOj2@ep-shrill-base-66245357.ap-southeast-1.postgres.vercel-storage.com:5432/verceldb"

    dsn := "host=ep-shrill-base-66245357-pooler.ap-southeast-1.postgres.vercel-storage.com user=default password=xBTaUne1kOj2 dbname=verceldb port=5432 sslmode=disable TimeZone=ap-southeast-1"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if(err != nil){
        panic("failed to connect database")
    }

    DB = db
}