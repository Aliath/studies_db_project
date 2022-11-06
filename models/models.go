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

	Season Season `gorm:"foreignKey:SeasonID"`
}

type Season struct {
	ID       int
	SeriesID int
	Number   int
	Name     string

	Series Series `gorm:"foreignKey:SeriesID"`
}

type Series struct {
	ID          int
	Name        string
	Description string
	GenreID     int

	Genre Genre `gorm:"foreignKey:GenreID"`
}

type EpisodeReview struct {
	ID         int
	EpisodeID  int
	ReviewerID int
	Review     string
	Rate       int
	CreatedAt  time.Time

	Episode  Episode `gorm:"foreignKey:EpisodeID"`
	Reviewer User    `gorm:"foreignKey:ReviewerID"`
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

	Actor   Actor   `gorm:"foreignKey:ActorID"`
	Episode Episode `gorm:"foreignKey:EpisodeID"`
}
