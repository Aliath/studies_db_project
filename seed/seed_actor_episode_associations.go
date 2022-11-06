package seed

import (
	"log"

	"github.com/Aliath/studies_db_project/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const (
	actorEpisodesAssociationsToSeed = episodesToSeed * 4
)

func SeedActorAssocations(db *gorm.DB) {
	log.Print("seeding actors <-> episodes associations...")

	associationsToAdd := make([]models.ActorsEpisodesAssociation, actorEpisodesAssociationsToSeed)

	// make sure that at least one actor is assigned
	for i := 0; i < episodesToSeed; i++ {
		episodeId := i + 1

		associationsToAdd[i].EpisodeID = episodeId
		associationsToAdd[i].ActorID = gofakeit.IntRange(1, actorsToSeed)
	}

	for i := episodesToSeed; i < actorEpisodesAssociationsToSeed; i++ {
		associationsToAdd[i].EpisodeID = gofakeit.IntRange(1, episodesToSeed)
		associationsToAdd[i].ActorID = gofakeit.IntRange(1, actorsToSeed)
	}

	db.CreateInBatches(associationsToAdd, operationsPerBatch)
	log.Print("actors <-> episodes associations have been seeded")
}
