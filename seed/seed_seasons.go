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
	seasonsToSeed = 200
)

func FillSeason(model *models.Season) {
	model.Name = ""

	if gofakeit.Bool() {
		model.Name += gofakeit.Adjective() + " "
	}

	model.Name += gofakeit.Verb()
	model.Name = cases.Title(language.English).String(model.Name)

	// skipping series_id and number for now
}

func SeedSeasons(db *gorm.DB) {
	log.Print("seeding seasons...")

	seasonsToAdd := make([]models.Season, seasonsToSeed)
	numberBySeries := make(map[int]int, 0)

	// make sure that every single series has it own season (at least 1)
	for seriesId := 0; seriesId < seriesToSeed; seriesId++ {
		FillSeason(&seasonsToAdd[seriesId])

		seasonsToAdd[seriesId].SeriesID = seriesId + 1

		numberBySeries[seriesId]++
		seasonsToAdd[seriesId].Number = numberBySeries[seriesId]
	}

	for i := seriesToSeed; i < seasonsToSeed; i++ {
		FillSeason(&seasonsToAdd[i])

		seasonsToAdd[i].SeriesID = gofakeit.IntRange(1, seriesToSeed)

		numberBySeries[seasonsToAdd[i].SeriesID]++
		seasonsToAdd[i].Number = numberBySeries[seasonsToAdd[i].SeriesID]
	}

	db.CreateInBatches(seasonsToAdd, operationsPerBatch)
	log.Print("seasons have been added successfully")
}
