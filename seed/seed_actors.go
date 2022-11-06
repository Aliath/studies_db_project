package seed

import (
	"log"
	"time"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const (
	actorsToSeed = 100
)

func SeedActors(db *gorm.DB) {
	log.Print("seeding actors...")

	actorsToAdd := make([]models.Actor, actorsToSeed)

	dateFrom, _ := time.Parse("2006-01-02", "1980-01-01")
	dateTo, _ := time.Parse("2006-01-02", "2000-01-01")

	for i := 0; i < actorsToSeed; i++ {
		actorsToAdd[i].DateOfBirth = gofakeit.DateRange(dateFrom, dateTo)
		actorsToAdd[i].Fullname = gofakeit.Name()
	}

	db.CreateInBatches(actorsToAdd, operationsPerBatch)
	log.Print("actors have been added successfully")
}
