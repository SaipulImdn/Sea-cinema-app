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

// ... Implement HTTP handlers for different routes and API endpoints

func (s *Server) Start(port int) {
	http.HandleFunc("/movies", s.handleGetMovies)
	http.HandleFunc("/movies/{id}", s.handleGetMovieDetails)
	// ... Define additional routes and handlers

	addr := ":" + strconv.Itoa(port)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
