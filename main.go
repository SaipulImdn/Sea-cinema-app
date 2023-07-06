package main

import (
	"github.com/SaipuImdn/sea-cinema-app/server"
	"github.com/SaipulImdn/sea-cinema-app/db"
)

func main() {
	database := db.NewDatabase()
	appServer := server.NewServer(database)
	appServer.Start(8080) // Replace with your desired port
}
