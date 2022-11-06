package seed

import (
	"log"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const (
	genresToSeed = 5
)

func SeedGenres(db *gorm.DB) {
	log.Print("seeding genres...")

	genreNames := [genresToSeed]string{
		"Horror",
		"Comedy",
		"Science Fiction",
		"Thriller",
		"Animated",
	}

	genresToAdd := make([]models.Genre, genresToSeed)

	for i := 0; i < genresToSeed; i++ {
		genresToAdd[i].Name = genreNames[i]
		genresToAdd[i].Description = gofakeit.LoremIpsumSentence(gofakeit.IntRange(8, 20))
	}

	db.CreateInBatches(genresToAdd, operationsPerBatch)
	log.Print("genres have been added successfully")
}
