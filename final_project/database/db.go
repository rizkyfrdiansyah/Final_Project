package database

import (
	"finalproject/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
    host     = "localhost"
    user     = "root"
    password = "Swadaya05@"
    dbPort   = "3306"
    dbName   = "mygramm"
    db       *gorm.DB
    err      error
)

func StartDB() {
    // Perbaikan pada format DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbName)
    
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Error connecting to database :", err)
    }

    fmt.Println("Connection success")
    db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
    return db
}
