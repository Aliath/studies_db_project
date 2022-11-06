package seed

import (
	"log"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

const (
	seriesToSeed = 100
)

func SeedSeries(db *gorm.DB) {
	log.Print("seeding series...")

	seriesToAdd := make([]models.Series, seriesToSeed)
	usedNames := make(map[string]bool)

	for i := 0; i < seriesToSeed; i++ {

		seriesToAdd[i].Name = ""

		if gofakeit.Bool() {
			seriesToAdd[i].Name += (gofakeit.VerbTransitive() + " ")
		}

		seriesToAdd[i].Name += gofakeit.VerbAction()
		seriesToAdd[i].Name = cases.Title(language.English).String(seriesToAdd[i].Name)

		if usedNames[seriesToAdd[i].Name] {
			i--
			continue
		}

		usedNames[seriesToAdd[i].Name] = true

		seriesToAdd[i].Description = gofakeit.LoremIpsumSentence(gofakeit.IntRange(5, 20))
		seriesToAdd[i].GenreID = gofakeit.IntRange(1, genresToSeed)
	}

	db.CreateInBatches(seriesToAdd, operationsPerBatch)
	log.Print("series have been added successfully")
}
