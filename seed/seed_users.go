package seed

import (
	"log"
	"time"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const (
	usersToSeed = 100
)

func SeedUsers(db *gorm.DB) {
	log.Print("seeding users...")

	usersToAdd := make([]models.User, usersToSeed)

	dateFrom, _ := time.Parse("2006-01-02", "2020-01-01")
	dateTo, _ := time.Parse("2006-01-02", "2022-10-01")

	for i := 0; i < actorsToSeed; i++ {
		usersToAdd[i].Email = gofakeit.Email()
		usersToAdd[i].FullName = gofakeit.Name()
		usersToAdd[i].PasswordHash = gofakeit.HexUint256()
		usersToAdd[i].PasswordSalt = gofakeit.HexUint256()
		usersToAdd[i].CreatedAt = gofakeit.DateRange(dateFrom, dateTo)

		usersToAdd[i].IsAdmin = i%2137 == 0
	}

	db.CreateInBatches(usersToAdd, operationsPerBatch)
	log.Print("users have been added successfully")
}
