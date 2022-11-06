package seed

import (
	"log"
	"time"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

const (
	episodesToSeed = 2000
)

func FillEpisode(model *models.Episode) {
	dateFrom, _ := time.Parse("2006-01-02", "2000-01-01")
	dateTo, _ := time.Parse("2006-01-02", "2022-10-01")

	model.PublicationDate = gofakeit.DateRange(dateFrom, dateTo)
	model.SeasonID = gofakeit.IntRange(1, seasonsToSeed)
	model.Description = gofakeit.LoremIpsumSentence(32)
	model.Duration = gofakeit.IntRange(60*20, 60*60)
	model.Name = gofakeit.Adjective() + " " + gofakeit.Verb() + " " + gofakeit.Noun()

	model.Name = cases.Title(language.English).String(model.Name)
}

func SeedEpisodes(db *gorm.DB) {
	log.Print("seeding episodes...")

	episodesToAdd := make([]models.Episode, episodesToSeed)
	numberByEpisodes := make(map[int]int, 0)

	// make sure that every single series has it own season (at least 1)
	for seasonIndex := 0; seasonIndex < seasonsToSeed; seasonIndex++ {
		FillEpisode(&episodesToAdd[seasonIndex])

		seasonId := seasonIndex + 1
		episodesToAdd[seasonIndex].SeasonID = seasonId

		numberByEpisodes[seasonId] = numberByEpisodes[seasonId] + 1

		episodesToAdd[seasonIndex].Number = numberByEpisodes[seasonId]
	}

	for i := seasonsToSeed; i < episodesToSeed; i++ {
		FillEpisode(&episodesToAdd[i])

		seasonId := gofakeit.IntRange(1, seasonsToSeed)
		episodesToAdd[i].SeasonID = seasonId

		numberByEpisodes[seasonId] = numberByEpisodes[seasonId] + 1
		episodesToAdd[i].Number = numberByEpisodes[seasonId]
	}

	db.CreateInBatches(episodesToAdd, operationsPerBatch)
	log.Print("episodes have been added successfully")
}
