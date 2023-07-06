package booking

import (
	"errors"

	"github.com/SaipulImdn/sea-cinema-app/db"
)

// BookingService provides movie ticket booking functionalities
type BookingService struct {
	db *db.Database
}

// NewBookingService creates a new instance of the booking service
func NewBookingService(database *db.Database) *BookingService {
	return &BookingService{
		db: database,
	}
}

// GetMovies returns a list of all available movies
func (b *BookingService) GetMovies() ([]db.Movie, error) {
	// Implement logic to retrieve movies from the database
	movies := b.db.GetMovies()
	if len(movies) == 0 {
		return nil, errors.New("no movies available")
	}

	return movies, nil
}

// GetMovieDetails returns the details of a specific movie
func (b *BookingService) GetMovieDetails(movieID int) (*db.Movie, error) {
	// Implement logic to retrieve movie details from the database
	movie := b.db.FindMovieByID(movieID)
	if movie == nil {
		return nil, errors.New("movie not found")
	}

	return movie, nil
}

// BookTicket allows a user to book a movie ticket
func (b *BookingService) BookTicket(userID, movieID, seatNumber int) error {
	// Implement logic to book a movie ticket
	user := b.db.FindUserByID(userID)
	if user == nil {
		return errors.New("user not found")
	}

	movie := b.db.FindMovieByID(movieID)
	if movie == nil {
		return errors.New("movie not found")
	}

	// Check if the seat is available
	if !b.db.IsSeatAvailable(movieID, seatNumber) {
		return errors.New("seat is not available")
	}

	// Check if the user's age meets the movie's age rating
	if user.Age < movie.AgeRating {
		return errors.New("user age does not meet the movie's age rating")
	}

	// Book the ticket and update the database
	err := b.db.BookTicket(userID, movieID, seatNumber)
	if err != nil {
		return err
	}

	return nil
}

// ... Implement additional ticket booking methods as needed
