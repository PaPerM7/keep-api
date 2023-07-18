package database

import (
	"fmt"
	"log"

	"keep-api/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {

	fmt.Print("db url: ", viper.Get("DB_HOST"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", viper.Get("DB_HOST"), viper.Get("DB_USERNAME"), viper.Get("DB_PASSWORD"), viper.Get("DB_NAME"), viper.Get("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Note{})

	DB = Dbinstance{
		Db: db,
	}
}
