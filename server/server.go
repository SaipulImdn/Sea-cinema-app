package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SaipulImdn/sea-cinema-app/auth"
	"github.com/SaipulImdn/sea-cinema-app/booking"
	"github.com/SaipulImdn/sea-cinema-app/db"
)

type Server struct {
	authService    *auth.AuthService
	bookingService *booking.BookingService
}

func NewServer(database *db.Database) *Server {
	authService := auth.NewAuthService(database)
	bookingService := booking.NewBookingService(database)

	return &Server{
		authService:    authService,
		bookingService: bookingService,
	}
}

// HandleGetMovies handles the "/movies" route to get a list of all available movies
func (s *Server) handleGetMovies(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get a list of all available movies from the booking service
	movies, err := s.bookingService.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the movies to JSON and write the response
	// You may need to use an encoding/json package or a framework for JSON serialization
	// Example: json.NewEncoder(w).Encode(movies)
}

// HandleGetMovieDetails handles the "/movies/{id}" route to get details of a specific movie
func (s *Server) handleGetMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Parse the movie ID from the request URL
	// You can use a router or the "github.com/gorilla/mux" package for URL routing
	// Example: movieID, err := strconv.Atoi(mux.Vars(r)["id"])
	// Handle any errors in parsing the ID

	// Implement logic to get the movie details from the booking service based on the movie ID
	movie, err := s.bookingService.GetMovieDetails(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the movie details to JSON and write the response
	// Example: json.NewEncoder(w).Encode(movie)
}

// ... Implement additional HTTP handlers for different routes and API endpoints

func (s *Server) Start(port int) {
	http.HandleFunc("/movies", s.handleGetMovies)
	http.HandleFunc("/movies/{id}", s.handleGetMovieDetails)
	// ... Define additional routes and handlers

	addr := ":" + strconv.Itoa(port)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
