package main

import (
	"log"
	"os"
	"time"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	operationsPerBatch = 100

	actorsToSeed                    = 100
	usersToSeed                     = 100
	seriesToSeed                    = 100
	episodesToSeed                  = 10000
	seasonsToSeed                   = 200
	genresToSeed                    = 5
	actorEpisodesAssociationsToSeed = 4
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
		usersToAdd[i].IsAdmin = gofakeit.Bool()
	}

	db.CreateInBatches(usersToAdd, operationsPerBatch)
	log.Print("users have been added successfully")
}

func SeedSeries(db *gorm.DB) {
	log.Print("seeding series...")

	seriesToAdd := make([]models.Series, seriesToSeed)

	for i := 0; i < seriesToSeed; i++ {
		seriesToAdd[i].Name = gofakeit.VerbTransitive() + " " + gofakeit.VerbAction()
		seriesToAdd[i].Description = gofakeit.LoremIpsumSentence(12)
		seriesToAdd[i].GenreId = gofakeit.IntRange(1, genresToSeed)
	}

	db.CreateInBatches(seriesToAdd, operationsPerBatch)
	log.Print("series have been added successfully")
}

func SeedGenres(db *gorm.DB) {
	log.Print("seeding genres...")

	genresToAdd := make([]models.Genre, genresToSeed)

	for i := 0; i < genresToSeed; i++ {
		genresToAdd[i].Name = gofakeit.Adjective()
		genresToAdd[i].Description = gofakeit.LoremIpsumSentence(10)
	}

	db.CreateInBatches(genresToAdd, operationsPerBatch)
	log.Print("genres have been added successfully")
}

func SeedSeasons(db *gorm.DB) {
	log.Print("seeding seasons...")

	seasonsToAdd := make([]models.Season, seasonsToSeed)
	numberBySeries := make(map[int]int, 0)

	for i := 0; i < seasonsToSeed; i++ {
		seasonsToAdd[i].Name = gofakeit.Adjective() + " " + gofakeit.Verb()
		seasonsToAdd[i].SeriesID = gofakeit.IntRange(1, seriesToSeed)

		numberBySeries[seasonsToAdd[i].SeriesID] += 1

		seasonsToAdd[i].Number = numberBySeries[seasonsToAdd[i].SeriesID]
	}

	db.CreateInBatches(seasonsToAdd, operationsPerBatch)
	log.Print("seasons have been added successfully")
}

func SeedEpisodes(db *gorm.DB) {
	log.Print("seeding episodes...")

	episodesToAdd := make([]models.Episode, episodesToSeed)
	numberByEpisodes := make(map[int]int, 0)

	dateFrom, _ := time.Parse("2006-01-02", "2000-01-01")
	dateTo, _ := time.Parse("2006-01-02", "2022-10-01")

	for i := 0; i < episodesToSeed; i++ {
		episodesToAdd[i].PublicationDate = gofakeit.DateRange(dateFrom, dateTo)
		episodesToAdd[i].SeasonID = gofakeit.IntRange(1, seasonsToSeed)
		episodesToAdd[i].Description = gofakeit.LoremIpsumSentence(32)
		episodesToAdd[i].Duration = gofakeit.IntRange(60*20, 60*60)
		episodesToAdd[i].Name = gofakeit.Adjective() + " " + gofakeit.Verb() + " " + gofakeit.Noun()

		numberByEpisodes[episodesToAdd[i].SeasonID] += 1
		episodesToAdd[i].Number = numberByEpisodes[episodesToAdd[i].SeasonID]
	}

	db.CreateInBatches(episodesToAdd, operationsPerBatch)
	log.Print("episodes have been added successfully")
}

func SeedActorAssocations(db *gorm.DB) {
	log.Print("seeding actors <-> episodes associations...")

	associationsToAdd := make([]models.ActorsEpisodesAssociation, 0)

	episodeIdentifiers := make([]int, episodesToSeed)
	for i := 1; i < episodesToSeed; i++ {
		episodeIdentifiers[i] = i
	}

	for actorId := 1; actorId <= actorsToSeed; actorId++ {
		alreadyAssigned := make(map[int]bool, 0)
		for i := 0; i < gofakeit.IntRange(1, 10); i++ {
			episodeId := gofakeit.RandomInt(episodeIdentifiers)

			if alreadyAssigned[episodeId] {
				continue
			}

			alreadyAssigned[episodeId] = true
			associationsToAdd = append(associationsToAdd, models.ActorsEpisodesAssociation{
				ActorID:   actorId,
				EpisodeID: episodeId,
			})
		}
	}

	db.CreateInBatches(associationsToAdd, operationsPerBatch)
	log.Print("actors <-> episodes associations have been seeded")
}

func main() {
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

	SeedActors(db)
	SeedUsers(db)
	SeedGenres(db)
	SeedSeries(db)
	SeedSeasons(db)
	SeedEpisodes(db)
	SeedActorAssocations(db)
}
