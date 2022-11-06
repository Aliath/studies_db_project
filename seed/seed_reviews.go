package seed

import (
	"log"
	"time"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const (
	reviewsToSeed = 500
)

func FillReview(model *models.EpisodeReview) {
	dateFrom, _ := time.Parse("2006-01-02", "2020-01-01")
	dateTo, _ := time.Parse("2006-01-02", "2022-10-01")

	model.CreatedAt = gofakeit.DateRange(dateFrom, dateTo)
	model.EpisodeID = gofakeit.IntRange(1, episodesToSeed)
	model.Rate = gofakeit.IntRange(1, 10)
	model.Review = gofakeit.LoremIpsumSentence(gofakeit.IntRange(8, 100))
	model.ReviewerID = gofakeit.IntRange(1, usersToSeed)
}

func SeedReviews(db *gorm.DB) {
	log.Print("seeding reviews...")

	reviewsToAdd := make([]models.EpisodeReview, reviewsToSeed)

	for i := 0; i < reviewsToSeed; i++ {
		FillReview(&reviewsToAdd[i])
	}

	db.CreateInBatches(reviewsToAdd, operationsPerBatch)
	log.Print("reviews have been added successfully")
}
