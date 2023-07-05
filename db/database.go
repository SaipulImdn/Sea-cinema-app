package db

type Movie struct {
	ID          int
	Title       string
	Description string
	AgeRating   int
	PosterURL   string
	TicketPrice float64
}

type User struct {
	ID         int
	UserID     int
	MovieiID   int
	SeatNumber int
}
