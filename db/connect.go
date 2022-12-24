package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/raihanbachdim/ProjectSIU/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Tidak bisa mengunduh file env nya bang")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak bisa menyambung ke database")
	} else {
		log.Println("Berhasil tersambung ke database")
	}
	DB = database
	database.AutoMigrate(
		&model.User{},
		&model.Stores{},
	)

}
