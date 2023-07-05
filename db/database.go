package db

import "sync"

type Movie struct {
	ID          int
	Title       string
	Description string
	AgeRating   int
	PosterURL   string
	TicketPrice float64
}

type User struct {
	ID       int
	Username string
	Password string
	Name     string
	Age      int
	Balance  float64
}

type Ticket struct {
	ID         int
	UserID     int
	MovieiID   int
	SeatNumber int
}

type Database struct {
	movies       []Movie
	users        []User
	tickets      []Ticket
	nextMovieID  int
	nextUserID   int
	nextTicketID int
	mu           sync.Mutex
}

func NewDatabse() *Database {
	return &Database{}
}
