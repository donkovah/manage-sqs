package config

import (
	"be/src/domain/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sakirsensoy/genv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type appConfig struct {
	Port       int
	Debug      bool
	DBPort     int
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

var App *appConfig
var DB *gorm.DB

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	App = &appConfig{
		Port:       genv.Key("PORT").Default(8080).Int(),
		Debug:      genv.Key("DEBUG").Default(false).Bool(),
		DBPort:     genv.Key("DB_PORT").Default(5432).Int(),
		DBHost:     genv.Key("DB_HOST").Default("localhost").String(),
		DBUser:     genv.Key("DB_USER").Default("gorm").String(),
		DBName:     genv.Key("DB_NAME").Default("gorm").String(),
		DBPassword: genv.Key("DB_PASSWORD").String(),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", App.DBHost, App.DBUser, App.DBPassword, App.DBName, App.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Task{}, &models.Project{}, &models.TaskLog{}, &models.Note{})
	DB = db
}
