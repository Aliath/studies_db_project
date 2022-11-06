package database

import (
	"log"
	"os"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	gofakeit.Seed(1)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresURL := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect with database", err)
	}

	db.AutoMigrate(&models.Actor{}, &models.ActorsEpisodesAssociation{}, &models.Episode{}, &models.EpisodeReview{}, &models.Genre{}, &models.Season{}, &models.Series{}, &models.User{})

	return db
}
