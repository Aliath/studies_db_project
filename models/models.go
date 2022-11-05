package models

import "time"

type Actor struct {
	ID          int
	Fullname    string
	DateOfBirth time.Time
}

type Episode struct {
	ID              int
	SeasonID        int
	Name            string
	Number          int
	Description     string
	Duration        int
	PublicationDate time.Time
}

type Season struct {
	ID       int
	SeriesID int
	Number   int
	Name     string
}

type Series struct {
	ID          int
	Name        string
	Description string
	GenreId     int
}

type EpisodeReview struct {
	ID         int
	EpisodeID  int
	ReviewerID int
	Review     string
	Rate       int
	CreatedAt  time.Time
}

type User struct {
	ID           int
	Email        string
	FullName     string
	PasswordHash string
	PasswordSalt string
	CreatedAt    time.Time
	IsAdmin      bool
}

type Genre struct {
	ID          int
	Name        string
	Description string
}

type ActorsEpisodesAssociation struct {
	ActorID   int
	EpisodeID int
}
