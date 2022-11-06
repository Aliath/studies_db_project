package main

import (
	"github.com/Aliath/studies_db_project/database"
	"github.com/Aliath/studies_db_project/seed"
)

func main() {
	db := database.GetDatabase()

	seed.SeedActors(db)
	seed.SeedUsers(db)
	seed.SeedGenres(db)
	seed.SeedSeries(db)
	seed.SeedSeasons(db)
	seed.SeedEpisodes(db)
	seed.SeedActorAssocations(db)
	seed.SeedReviews(db)
}
